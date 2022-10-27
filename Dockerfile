################
# BUILD BINARY #
################

FROM golang:1.18.2-alpine3.15 as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.Version=v1.0.0'" .

#####################
# MAKE SMALL BINARY #
#####################
FROM scratch

# Copy the executable.
WORKDIR /app

COPY --from=builder /app/demogogen1 /usr/bin/
COPY --from=builder /app/config.json /app

# ENTRYPOINT ["demogogen1", "myapp"]
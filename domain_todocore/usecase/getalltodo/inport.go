package getalltodo

import (
	"demogogen1/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	Page int64
	Size int64
}

type InportResponse struct {
	Count int64
	Items []any
}

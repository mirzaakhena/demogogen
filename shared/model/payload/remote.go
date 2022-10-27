package payload

import "demogogen1/shared/gogen"

type Args struct {
	Type      string                `json:"type"`
	Data      any                   `json:"data"`
	Publisher gogen.ApplicationData `json:"publisher"`
	TraceID   string                `json:"traceId"`
}

type Reply struct {
	Success      bool
	ErrorMessage string
	Data         any
}

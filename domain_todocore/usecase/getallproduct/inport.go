package getallproduct

import (
	"demogogen1/domain_todocore/model/repository"
	"demogogen1/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	repository.FindAllProductFilterRequest
}

type InportResponse struct {
	Count int64
	Items []any
}

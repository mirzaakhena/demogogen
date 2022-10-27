package runtodocreate

import (
	"demogogen1/domain_todocore/model/entity"
	"demogogen1/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	entity.TodoCreateRequest
}

type InportResponse struct {
}

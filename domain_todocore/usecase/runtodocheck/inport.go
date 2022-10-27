package runtodocheck

import (
	"demogogen1/domain_todocore/model/vo"
	"demogogen1/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	TodoID vo.TodoID
}

type InportResponse struct {
}

package getoneproduct

import (
	"demogogen1/domain_todocore/model/entity"
	"demogogen1/domain_todocore/model/vo"
	"demogogen1/shared/gogen"
)

type Inport gogen.Inport[InportRequest, InportResponse]

type InportRequest struct {
	ProductID vo.ProductID
}

type InportResponse struct {
	Product *entity.Product
}

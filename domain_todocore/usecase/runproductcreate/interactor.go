package runproductcreate

import (
	"context"
	"demogogen1/domain_todocore/model/entity"
)

//go:generate mockery --name Outport -output mocks/

type productCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &productCreateInteractor{
		outport: outputPort,
	}
}

func (r *productCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	productObj, err := entity.NewProduct(req.ProductCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveProduct(ctx, productObj)
	if err != nil {
		return nil, err
	}

	res.ProductID = productObj.ID

	return res, nil
}

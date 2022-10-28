package getoneproduct

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type productGetOneInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &productGetOneInteractor{
		outport: outputPort,
	}
}

func (r *productGetOneInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	productObj, err := r.outport.FindOneProductByID(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}

	if productObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	res.Product = productObj

	return res, nil
}

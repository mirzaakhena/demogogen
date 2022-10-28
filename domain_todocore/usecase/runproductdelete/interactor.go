package runproductdelete

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type productDeleteInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &productDeleteInteractor{
		outport: outputPort,
	}
}

func (r *productDeleteInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	productObj, err := r.outport.FindOneProductByID(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}
	if productObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	// you may need to check the authorization part here
	// is this user allowed to perform this action ?

	err = r.outport.DeleteProduct(ctx, req.ProductID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

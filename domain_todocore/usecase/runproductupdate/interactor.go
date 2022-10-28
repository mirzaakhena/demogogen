package runproductupdate

import (
	"context"
	"fmt"
)

//go:generate mockery --name Outport -output mocks/

type productUpdateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &productUpdateInteractor{
		outport: outputPort,
	}
}

func (r *productUpdateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

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

	err = productObj.Update(req.ProductUpdateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveProduct(ctx, productObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}

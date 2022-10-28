package getallproduct

import (
	"context"
	"demogogen1/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type productGetAllInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &productGetAllInteractor{
		outport: outputPort,
	}
}

func (r *productGetAllInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	productObj, count, err := r.outport.FindAllProduct(ctx, req.FindAllProductFilterRequest)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(productObj)

	return res, nil
}

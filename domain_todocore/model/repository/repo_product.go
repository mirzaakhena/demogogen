package repository

import (
	"context"
	"demogogen1/domain_todocore/model/entity"
	"demogogen1/domain_todocore/model/vo"
)

type SaveProductRepo interface {
	SaveProduct(ctx context.Context, obj *entity.Product) error
}

type FindAllProductFilterRequest struct {
	Page int64
	Size int64
	// add other field to filter here ...
}

type FindAllProductRepo interface {
	FindAllProduct(ctx context.Context, req FindAllProductFilterRequest) ([]*entity.Product, int64, error)
}

type DeleteProductRepo interface {
	DeleteProduct(ctx context.Context, productID vo.ProductID) error
}

type FindOneProductByIDRepo interface {
	FindOneProductByID(ctx context.Context, productID vo.ProductID) (*entity.Product, error)
}

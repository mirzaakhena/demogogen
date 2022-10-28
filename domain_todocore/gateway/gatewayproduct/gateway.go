package gatewayproduct

import (
	"context"
	"demogogen1/domain_todocore/model/entity"
	"demogogen1/domain_todocore/model/repository"
	"demogogen1/domain_todocore/model/vo"
	"demogogen1/shared/gogen"
	"demogogen1/shared/infrastructure/config"
	"demogogen1/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
	}
}

func (r *gateway) FindAllProduct(ctx context.Context, req repository.FindAllProductFilterRequest) ([]*entity.Product, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	return nil, 0, nil
}

func (r *gateway) FindOneProductByID(ctx context.Context, productID vo.ProductID) (*entity.Product, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveProduct(ctx context.Context, obj *entity.Product) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) DeleteProduct(ctx context.Context, productID vo.ProductID) error {
	r.log.Info(ctx, "called")

	return nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	return nil, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	return nil
}

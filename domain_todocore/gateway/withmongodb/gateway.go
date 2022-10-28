package withmongodb

import (
	"context"
	"demogogen1/domain_todocore/model/entity"
	"demogogen1/domain_todocore/model/vo"
	"demogogen1/shared/gogen"
	"demogogen1/shared/infrastructure/config"
	"demogogen1/shared/infrastructure/database"
	"demogogen1/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config

	repo database.Repository[entity.Todo]
}

// NewGateway ...
func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	db := database.NewDatabase("todo_db")

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		repo:    database.NewMongoGateway[entity.Todo](db),
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	param := database.NewDefaultParam().
		SetPage(page).
		SetSize(size)

	res := make([]*entity.Todo, 0)
	count, err := r.repo.GetAll(param, &res)
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (r *gateway) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	filter := map[string]any{
		"_id": todoID,
	}

	var res entity.Todo
	err := r.repo.GetOne(filter, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called %v", obj)

	err := r.repo.InsertOrUpdate(obj)
	if err != nil {
		return err
	}

	return nil
}

package runtodocreate

import (
	"context"
	"demogogen1/domain_todocore/model/entity"
)

type runTodoCreateInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCreateInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCreateInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := entity.NewTodo(req.TodoCreateRequest)
	if err != nil {
		return nil, err
	}

	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}

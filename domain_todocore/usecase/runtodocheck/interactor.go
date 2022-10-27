package runtodocheck

import (
	"context"
	"fmt"
)

type runTodoCheckInteractor struct {
	outport Outport
}

func NewUsecase(outputPort Outport) Inport {
	return &runTodoCheckInteractor{
		outport: outputPort,
	}
}

func (r *runTodoCheckInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObj, err := r.outport.FindOneTodoByID(ctx, req.TodoID)
	if err != nil {
		return nil, err
	}
	if todoObj == nil {
		return nil, fmt.Errorf("object not found")
	}

	err = r.outport.SaveTodo(ctx, todoObj)
	if err != nil {
		return nil, err
	}

	return res, nil
}

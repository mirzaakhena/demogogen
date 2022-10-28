package entity

import (
	"demogogen1/domain_todocore/model/errorenum"
	"demogogen1/domain_todocore/model/vo"
	"time"
)

type Todo struct {
	ID      vo.TodoID `bson:"_id" json:"id"`
	Message string    `bson:"message" json:"message"`
	Checked bool      `bson:"checked" json:"checked"`
	Created time.Time `bson:"created" json:"created"`
}

type TodoCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
	Message      string    `json:"message"`
}

func NewTodo(req TodoCreateRequest) (*Todo, error) {

	id, err := vo.NewTodoID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	if req.Message == "" {
		return nil, errorenum.MessageMustNotEmpty
	}

	var obj Todo
	obj.ID = id
	obj.Message = req.Message
	obj.Checked = false
	obj.Created = req.Now

	return &obj, nil
}

func (r *Todo) Check() error {

	if r.Checked {
		return errorenum.TodoAlreadyChecked.Var(r.Message)
	}

	r.Checked = true
	return nil
}

type TodoUpdateRequest struct {
	// add field to update here ...
}

func (r *Todo) Update(req TodoUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}

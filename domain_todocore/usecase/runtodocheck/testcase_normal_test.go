package runtodocheck

import (
	"context"
	"fmt"
	"testing"
	"time"

	"demogogen1/domain_todocore/model/entity"
	"demogogen1/domain_todocore/model/vo"
)

type mockOutportNormal struct {
	t *testing.T
}

// TestCaseNormal is for the case where ...
// explain the purpose of this test here with human readable naration...
func TestCaseNormal(t *testing.T) {

	ctx := context.Background()

	mockOutport := mockOutportNormal{
		t: t,
	}

	res, err := NewUsecase(&mockOutport).Execute(ctx, InportRequest{
		TodoID: "ABC",
	})

	if err != nil {
		t.Errorf("%v", err.Error())
		t.FailNow()
	}

	t.Logf("%v", res)

}

func (r *mockOutportNormal) FindOneTodoByID(ctx context.Context, todoID vo.TodoID) (*entity.Todo, error) {

	if todoID == "ABC" {
		return &entity.Todo{
			ID:      "ABC",
			Message: "This is a test",
			Checked: false,
			Created: time.Now(),
		}, nil
	}

	return nil, nil
}

func (r *mockOutportNormal) SaveTodo(ctx context.Context, obj *entity.Todo) error {

	fmt.Printf(">>>>> %v\n", obj)

	return nil
}

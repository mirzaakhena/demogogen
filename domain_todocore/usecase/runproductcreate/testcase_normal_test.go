package runproductcreate

import (
	"context"
	"testing"

	"demogogen1/domain_todocore/model/entity"
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

	res, err := NewUsecase(&mockOutport).Execute(ctx, InportRequest{})

	if err != nil {
		t.Errorf("%v", err.Error())
		t.FailNow()
	}

	t.Logf("%v", res)

}

func (r *mockOutportNormal) SaveProduct(ctx context.Context, obj *entity.Product) error {

	return nil
}

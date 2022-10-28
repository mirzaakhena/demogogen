package vo

import (
	"fmt"
	"time"
)

type ProductID string

func NewProductID(randomStringID string, now time.Time) (ProductID, error) {
	var obj = ProductID(fmt.Sprintf("OBJ-%s-%s", now.Format("060102"), randomStringID))
	return obj, nil
}

func (r ProductID) String() string {
	return string(r)
}

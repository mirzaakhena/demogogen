package entity

import (
	"demogogen1/domain_todocore/model/vo"
	"time"
)

type Product struct {
	ID      vo.ProductID `bson:"_id" json:"id"`
	Created time.Time    `bson:"created" json:"created"`
}

type ProductCreateRequest struct {
	RandomString string    `json:"-"`
	Now          time.Time `json:"-"`
}

func NewProduct(req ProductCreateRequest) (*Product, error) {

	id, err := vo.NewProductID(req.RandomString, req.Now)
	if err != nil {
		return nil, err
	}

	// add validation and assignment value here ...

	var obj Product
	obj.ID = id
	obj.Created = req.Now

	return &obj, nil
}

type ProductUpdateRequest struct {
	// add field to update here ...
}

func (r *Product) Update(req ProductUpdateRequest) error {

	// add validation and assignment value here ...

	return nil
}

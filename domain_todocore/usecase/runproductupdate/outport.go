package runproductupdate

import (
	"demogogen1/domain_todocore/model/repository"
)

type Outport interface {
	repository.FindOneProductByIDRepo
	repository.SaveProductRepo
}

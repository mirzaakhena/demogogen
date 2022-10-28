package runproductdelete

import (
	"demogogen1/domain_todocore/model/repository"
)

type Outport interface {
	repository.FindOneProductByIDRepo
	repository.DeleteProductRepo
}

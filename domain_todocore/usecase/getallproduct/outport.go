package getallproduct

import (
	"demogogen1/domain_todocore/model/repository"
)

type Outport interface {
	repository.FindAllProductRepo
}

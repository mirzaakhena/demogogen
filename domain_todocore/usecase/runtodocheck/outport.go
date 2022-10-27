package runtodocheck

import "demogogen1/domain_todocore/model/repository"

type Outport interface {
	repository.FindOneTodoByIDRepo
	repository.SaveTodoRepo
}

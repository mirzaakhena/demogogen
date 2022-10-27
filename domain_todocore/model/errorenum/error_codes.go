package errorenum

import "demogogen1/shared/model/apperror"

const (
	SomethingError     apperror.ErrorType = "ER0000 something error"
	TodoAlreadyChecked apperror.ErrorType = "ER0001 todo with the message \"%s\" already checked"
)

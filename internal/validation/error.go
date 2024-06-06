package validation

import (
	"errors"
	"fmt"
)

func NewInternalErr() error {
	return errors.New("internal error")
}

func NewNotFoundErr(entity string) error {
	return &ErrNotFound{
		Message: fmt.Sprintf("%s not found", entity),
	}
}

type ErrNotFound struct {
	Message string `json:"message"`
}

func (e *ErrNotFound) Error() string {
	return e.Message

}

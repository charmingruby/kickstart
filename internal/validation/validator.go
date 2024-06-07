package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validationErrs := err.(validator.ValidationErrors)
	validationErr := validationErrs[0]

	field := strings.ToLower(validationErr.StructField())
	switch validationErr.Tag() {
	case "required":
		return NewValidationErr(ErrRequired(field))
	case "max":
		return NewValidationErr(ErrMaxLength(field, validationErr.Param()))
	case "min":
		return NewValidationErr(ErrMinLength(field, validationErr.Param()))
	}

	return nil
}

func NewValidationErr(msg string) error {
	return &ErrInternal{
		Message: msg,
	}
}

type ErrValidation struct {
	Message string `json:"message"`
}

func (e *ErrValidation) Error() string {
	return e.Message
}

func ErrRequired(field string) string {
	return fmt.Sprintf("%s is required", field)
}

func ErrMinLength(field string, min string) string {
	return fmt.Sprintf("%s must have a minimum of %s", field, min)
}

func ErrMaxLength(field string, max string) string {
	return fmt.Sprintf("%s must have a maximum of %s", field, max)
}

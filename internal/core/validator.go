package core

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(obj interface{}) error {
	err := newValidator(obj)
	field := strings.ToLower(err.StructField())

	if err == nil {
		return nil
	}

	switch err.Tag() {
	case "required":
		return NewValidationErr(ErrRequired(field))
	case "max":
		return NewValidationErr(ErrMaxLength(field, err.Param()))
	case "min":
		return NewValidationErr(ErrMinLength(field, err.Param()))
	}

	return NewValidationErr(fmt.Sprintf("%s validation error on %s", field, err.Tag()))
}

func newValidator(obj interface{}) validator.FieldError {
	validate := validator.New()
	err := validate.Struct(obj)
	if err == nil {
		return nil
	}

	validationErrs := err.(validator.ValidationErrors)
	validationErr := validationErrs[0]

	return validationErr
}

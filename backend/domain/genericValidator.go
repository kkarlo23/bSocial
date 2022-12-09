package domain

import (
	"github.com/go-playground/validator/v10"
)

type validateInterface interface {
	User | Post | Comment | ApiRegister | ApiLogin
}

var validate = validator.New()

type ErrorValidate struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateType[T validateInterface](model T) []*ErrorValidate {
	var errors []*ErrorValidate

	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorValidate
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

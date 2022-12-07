package domain

import (
	"github.com/go-playground/validator/v10"
)

type validateInterface interface {
	User | Post | ApiRegister | ApiLogin
}

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateType[T validateInterface](model T) []*ErrorResponse {
	var errors []*ErrorResponse

	err := validate.Struct(model)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

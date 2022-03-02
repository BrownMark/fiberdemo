package utils

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorResponse struct {
	Param   string
	Message string
	Value   string
} //@name ErrorResponse

func messageForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error()
}

func ValidateStruct(item interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(item)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Param = err.Field()
			element.Message = messageForTag(err)
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

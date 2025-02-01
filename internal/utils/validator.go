package utils

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func InitValidator() {
	if Validate == nil {
		Validate = validator.New()
	}
}

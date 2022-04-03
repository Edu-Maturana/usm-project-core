package utils

import "github.com/go-playground/validator/v10"

func ValidateData(data interface{}) error {
	validator := validator.New()
	validationErrors := validator.Struct(data)
	if validationErrors != nil {
		return validationErrors
	}
	return nil
}

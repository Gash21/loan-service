package validator

import "github.com/go-playground/validator/v10"

type ValidatorService struct {
	Validate *validator.Validate
}
type IValidatorService interface {
	ValidateStruct(input interface{}) error
}

package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidatorService struct {
	Validate   *validator.Validate
	Translator ut.Translator
}
type ValidationError struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}
type IValidatorService interface {
	ValidateStruct(input interface{}) error
	TranslateError(err error) []ValidationError
}

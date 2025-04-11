package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func NewValidator() (IValidatorService, error) {
	v := validator.New()
	english := en.New()
	uni := ut.New(english)
	trans, _ := uni.GetTranslator("en")

	return &ValidatorService{
		Validate:   v,
		Translator: trans,
	}, nil
}

func (s ValidatorService) ValidateStruct(input interface{}) error {
	return s.Validate.Struct(input)
}

func (s *ValidatorService) TranslateError(err error) []ValidationError {
	// check if err is nil
	if err == nil {
		return nil
	}

	// check if err is validator.ValidationErrors
	validatorErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	// translate each error
	errors := []ValidationError{}
	for _, e := range validatorErrs {
		errors = append(errors, ValidationError{
			Field:   e.Field(),
			Value:   e.Value(),
			Message: e.Translate(s.Translator),
		})
	}

	return errors
}

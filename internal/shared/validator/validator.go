package validator

import "github.com/go-playground/validator/v10"

func NewValidator() (IValidatorService, error) {
	v := validator.New()

	return &ValidatorService{
		Validate: v,
	}, nil
}

func (s ValidatorService) ValidateStruct(input interface{}) error {
	return s.Validate.Struct(input)
}

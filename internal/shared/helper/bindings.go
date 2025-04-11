package helper

import (
	"net/http"

	"github.com/Gash21/amartha-test/internal/shared/logger"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

const (
	// Common error messages
	ContextName          = "Helper"
	ErrorValidatePayload = "Failed to validate payload"
)

type ModelValidationError struct {
	Code         int
	ResponseBody JSONResult
}

func (e *ModelValidationError) Error() string {
	return "Failed to validate request body"
}

func BindModel(log *zap.Logger, c *fiber.Ctx, m interface{}) {
	// create local logger
	l := logger.WithId(log, ContextName, "BindModel")

	// try binding request body
	if err := c.BodyParser(m); err != nil {
		l.Debug("Error when binding from body", zap.Error(err))
	}

	// try binding path parameters
	if err := c.QueryParser(m); err != nil {
		l.Debug("Error when binding from query string", zap.Error(err))
	}

	// try binding path parameters
	if err := c.ParamsParser(m); err != nil {
		l.Debug("Error when binding from path parameters", zap.Error(err))
	}

}

func ValidateModel(log *zap.Logger, v validator.IValidatorService, m interface{}) error {
	// create local logger
	l := logger.WithId(log, ContextName, "ValidateModel")

	// try validate model
	if err := v.ValidateStruct(m); err != nil {
		// log error
		l.Error(ErrorValidatePayload, zap.Error(err))

		// translate error
		localizedErr := v.TranslateError(err)

		// return error
		result := ResponseFailed(http.StatusBadRequest, ErrorValidatePayload, localizedErr)
		return &ModelValidationError{
			Code:         result.Code,
			ResponseBody: result,
		}
	}

	return nil
}

// BindAndValidateModel binds and validates a model using the provided logger, context, validator service, and model.
//
// Parameters:
//   - logger: the zap logger
//   - context: the fiber context
//   - validator: the validator service
//   - model: the model to bind and validate
//
// Return type: error
func BindAndValidateModel[T any](logger *zap.Logger, context *fiber.Ctx, validator validator.IValidatorService, model *T) *ModelValidationError {
	// Bind model
	BindModel(logger, context, model)

	// Validate model
	if err := ValidateModel(logger, validator, model); err != nil {
		perr := err.(*ModelValidationError)
		return perr
	}
	return nil
}

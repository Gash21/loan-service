package deps

import (
	"github.com/Gash21/amartha-test/internal/shared/config"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"github.com/Gash21/amartha-test/internal/shared/mailer"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Instance struct {
	Fiber     *fiber.App
	Config    config.GlobalConfig
	DB        *database.Database
	Validator validator.IValidatorService
	Logger    *zap.Logger
	Mailer    mailer.IMailerAdapter
}

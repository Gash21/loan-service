package deps

import (
	"github.com/Gash21/amartha-test/internal/shared/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Instance struct {
	Fiber  *fiber.App
	DB     *database.Database
	Logger *zap.Logger
}

type Dependency struct {
	DB     *database.Database
	Logger *zap.Logger
}

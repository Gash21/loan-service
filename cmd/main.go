package main

import (
	"fmt"

	"github.com/Gash21/amartha-test/internal/shared/config"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"github.com/Gash21/amartha-test/internal/shared/logger"
	"github.com/Gash21/amartha-test/pkg/deps"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}
	l := logger.NewLogger(cfg.ServiceName, cfg.LogLevel)

	dep := &deps.Dependency{
		DB: &database.Database{
			Gorm: nil,
		},
		Logger: l,
	}

	app := BootstrapApp(dep)
	l.Info("Starting HTTP Server", zap.Int("port", cfg.Port))
	if err := app.Fiber.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"

	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/domain/document"
	"github.com/Gash21/amartha-test/internal/domain/investor"
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/config"
	"github.com/Gash21/amartha-test/internal/shared/database"
	"github.com/Gash21/amartha-test/internal/shared/logger"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/Gash21/amartha-test/pkg/deps"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		panic(err)
	}
	l := logger.NewLogger(cfg.ServiceName, cfg.LogLevel)

	v, _ := validator.NewValidator()

	db, err := database.NewDatabase(database.SQLDatabaseOpts{
		GormConfig: gorm.Config{
			PrepareStmt: true,
		},
		Debug:  cfg.Debug,
		Logger: l,
		DBName: cfg.DBName,
	})

	if cfg.AutoMigrate {
		db.Gorm.AutoMigrate(&loan.Loan{}, &borrower.Borrower{}, &investor.Investor{}, &document.Document{})
	}

	if err != nil {
		panic(err)
	}

	dep := &deps.Instance{
		DB:        db,
		Logger:    l,
		Validator: v,
	}

	app := BootstrapApp(dep)
	l.Info("Starting HTTP Server", zap.Int("port", cfg.Port))
	if err := app.Fiber.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		panic(err)
	}
}

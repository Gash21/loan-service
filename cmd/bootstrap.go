package main

import (
	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/domain/document"
	"github.com/Gash21/amartha-test/internal/domain/employee"
	"github.com/Gash21/amartha-test/internal/domain/investor"
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/domain/loan_investor"
	borrowerHandler "github.com/Gash21/amartha-test/internal/presentation/http/handler/borrower"
	employeeHandler "github.com/Gash21/amartha-test/internal/presentation/http/handler/employee"
	loanHandler "github.com/Gash21/amartha-test/internal/presentation/http/handler/loan"
	"github.com/Gash21/amartha-test/pkg/deps"
	gojson "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func BootstrapApp(dep *deps.Instance) *deps.Instance {
	if dep.Config.AutoMigrate {
		dep.DB.Gorm.AutoMigrate(
			&loan.Loan{},
			&employee.Employee{},
			&borrower.Borrower{},
			&investor.Investor{},
			&document.Document{},
			&loan_investor.LoanInvestor{},
		)
	}

	e := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		JSONEncoder:           gojson.Marshal,
		JSONDecoder:           gojson.Unmarshal,
		BodyLimit:             11 * 1024 * 1024,
	})

	e.Use(cors.New())
	e.Use(recover.New())

	instance := &deps.Instance{
		Fiber:     e,
		Logger:    dep.Logger,
		Validator: dep.Validator,
		Config:    dep.Config,
		DB:        dep.DB,
	}

	loanHandler.RegisterAPI(instance)
	borrowerHandler.RegisterAPI(instance)
	employeeHandler.RegisterAPI(instance)

	return instance
}

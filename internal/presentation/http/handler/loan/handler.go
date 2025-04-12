package loan

import (
	"github.com/Gash21/amartha-test/internal/application/loan"
	borrowerRepository "github.com/Gash21/amartha-test/internal/infrastructure/borrower"
	employeeRepository "github.com/Gash21/amartha-test/internal/infrastructure/employee"
	investorRepository "github.com/Gash21/amartha-test/internal/infrastructure/investor"
	loanRepository "github.com/Gash21/amartha-test/internal/infrastructure/loan"
	"github.com/Gash21/amartha-test/internal/shared/config"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/Gash21/amartha-test/pkg/deps"
	"go.uber.org/zap"
)

const ContextName = "Presentation.Http.Loan"

type (
	Handler struct {
		Logger    *zap.Logger
		Config    config.GlobalConfig
		Validator validator.IValidatorService
		Usecase   loan.IUsecase
	}
)

func NewHandler(deps *deps.Instance) *Handler {
	return &Handler{
		Logger:    deps.Logger,
		Validator: deps.Validator,
		Config:    deps.Config,
		Usecase: loan.NewUsecase(loan.Usecase{
			Logger:             deps.Logger,
			Mailer:             deps.Mailer,
			Config:             deps.Config,
			InvestorRepository: investorRepository.NewRepository(deps.DB),
			BorrowerRepository: borrowerRepository.NewRepository(deps.DB),
			LoanRepository:     loanRepository.NewRepository(deps.DB),
			EmployeeRepository: employeeRepository.NewRepository(deps.DB),
		}),
	}
}

func RegisterAPI(d *deps.Instance) *Handler {
	h := NewHandler(d)

	g := d.Fiber.Group("/api/v1/loans")
	g.Get("/", h.List)
	g.Get("/:id", h.Detail)
	g.Post("/propose", h.Propose)
	g.Patch("/:id/approve", h.Approve)
	g.Patch("/:id/invest", h.Invest)
	g.Patch("/:id/disburse", h.Disburse)

	return h
}

package loan

import (
	"github.com/Gash21/amartha-test/internal/domain/loan"
	documentRepository "github.com/Gash21/amartha-test/internal/infrastructure/document"
	loanRepository "github.com/Gash21/amartha-test/internal/infrastructure/loan"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/Gash21/amartha-test/pkg/deps"
	"go.uber.org/zap"
)

const ContextName = "Presentation.Http.Loan"

type (
	Handler struct {
		Logger    *zap.Logger
		Validator validator.IValidatorService
		Usecase   loan.IUsecase
	}
)

func NewHandler(deps *deps.Instance) *Handler {
	return &Handler{
		Logger:    deps.Logger,
		Validator: deps.Validator,
		Usecase: loan.NewUsecase(loan.Usecase{
			Logger:             deps.Logger,
			DocumentRepository: documentRepository.NewRepository(deps.DB),
			LoanRepository:     loanRepository.NewRepository(deps.DB),
		}),
	}
}

func RegisterAPI(d *deps.Instance) *Handler {
	h := NewHandler(d)

	g := d.Fiber.Group("/api/v1/loans")

	g.Get("/", h.Propose)
	g.Post("/", h.Approved)
	g.Post("/", h.Invested)
	g.Post("/", h.Disbursed)

	return h
}

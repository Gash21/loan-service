package investor

import (
	"github.com/Gash21/amartha-test/internal/application/investor"
	investorRepository "github.com/Gash21/amartha-test/internal/infrastructure/investor"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/Gash21/amartha-test/pkg/deps"
	"go.uber.org/zap"
)

const ContextName = "Presentation.Http.Loan"

type (
	Handler struct {
		Logger    *zap.Logger
		Validator validator.IValidatorService
		Usecase   investor.IUsecase
	}
)

func NewHandler(deps *deps.Instance) *Handler {
	return &Handler{
		Logger:    deps.Logger,
		Validator: deps.Validator,
		Usecase: investor.NewUsecase(investor.Usecase{
			Logger:             deps.Logger,
			InvestorRepository: investorRepository.NewRepository(deps.DB),
		}),
	}
}

func RegisterAPI(d *deps.Instance) *Handler {
	h := NewHandler(d)

	g := d.Fiber.Group("/api/v1/investors")
	g.Post("/", h.Create)

	return h
}

package borrower

import (
	"github.com/Gash21/amartha-test/internal/application/borrower"
	borrowerRepository "github.com/Gash21/amartha-test/internal/infrastructure/borrower"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/Gash21/amartha-test/pkg/deps"
	"go.uber.org/zap"
)

const ContextName = "Presentation.Http.Loan"

type (
	Handler struct {
		Logger    *zap.Logger
		Validator validator.IValidatorService
		Usecase   borrower.IUsecase
	}
)

func NewHandler(deps *deps.Instance) *Handler {
	return &Handler{
		Logger:    deps.Logger,
		Validator: deps.Validator,
		Usecase: borrower.NewUsecase(borrower.Usecase{
			Logger:             deps.Logger,
			BorrowerRepository: borrowerRepository.NewRepository(deps.DB),
		}),
	}
}

func RegisterAPI(d *deps.Instance) *Handler {
	h := NewHandler(d)

	g := d.Fiber.Group("/api/v1/borrowers")
	g.Post("/", h.Create)

	return h
}

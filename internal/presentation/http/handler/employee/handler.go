package employee

import (
	"github.com/Gash21/amartha-test/internal/application/employee"
	employeeRepository "github.com/Gash21/amartha-test/internal/infrastructure/employee"
	"github.com/Gash21/amartha-test/internal/shared/validator"
	"github.com/Gash21/amartha-test/pkg/deps"
	"go.uber.org/zap"
)

const ContextName = "Presentation.Http.Loan"

type (
	Handler struct {
		Logger    *zap.Logger
		Validator validator.IValidatorService
		Usecase   employee.IUsecase
	}
)

func NewHandler(deps *deps.Instance) *Handler {
	return &Handler{
		Logger:    deps.Logger,
		Validator: deps.Validator,
		Usecase: employee.NewUsecase(employee.Usecase{
			Logger:             deps.Logger,
			EmployeeRepository: employeeRepository.NewRepository(deps.DB),
		}),
	}
}

func RegisterAPI(d *deps.Instance) *Handler {
	h := NewHandler(d)

	g := d.Fiber.Group("/api/v1/employees")
	g.Post("/", h.Create)

	return h
}

package employee

import (
	"context"

	"github.com/Gash21/amartha-test/internal/application/employee/dto"
	"github.com/Gash21/amartha-test/internal/domain/employee"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"go.uber.org/zap"
)

type (
	Usecase struct {
		EmployeeRepository employee.EmployeeRepository
		Logger             *zap.Logger
	}
	IUsecase interface {
		Create(context.Context, dto.CreateRequest) helper.JSONResult
	}
)

func NewUsecase(opts Usecase) IUsecase {
	return &Usecase{
		EmployeeRepository: opts.EmployeeRepository,
		Logger:             opts.Logger,
	}
}

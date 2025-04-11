package investor

import (
	"context"

	"github.com/Gash21/amartha-test/internal/application/investor/dto"
	"github.com/Gash21/amartha-test/internal/domain/investor"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"go.uber.org/zap"
)

type (
	Usecase struct {
		InvestorRepository investor.InvestorRepository
		Logger             *zap.Logger
	}
	IUsecase interface {
		Create(context.Context, dto.CreateRequest) helper.JSONResult
	}
)

func NewUsecase(opts Usecase) IUsecase {
	return &Usecase{
		InvestorRepository: opts.InvestorRepository,
		Logger:             opts.Logger,
	}
}

package borrower

import (
	"context"

	"github.com/Gash21/amartha-test/internal/application/borrower/dto"
	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"go.uber.org/zap"
)

type (
	Usecase struct {
		BorrowerRepository borrower.BorrowerRepository
		Logger             *zap.Logger
	}
	IUsecase interface {
		Create(context.Context, dto.CreateRequest) helper.JSONResult
	}
)

func NewUsecase(opts Usecase) IUsecase {
	return &Usecase{
		BorrowerRepository: opts.BorrowerRepository,
		Logger:             opts.Logger,
	}
}

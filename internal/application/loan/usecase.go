package loan

import (
	"github.com/Gash21/amartha-test/internal/domain/document"
	"github.com/Gash21/amartha-test/internal/domain/loan"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"go.uber.org/zap"
)

type (
	Usecase struct {
		LoanRepository     loan.LoanRepository
		DocumentRepository document.DocumentRepository
		Logger             *zap.Logger
	}
	IUsecase interface {
		Propose() helper.JSONResult
		Approve() helper.JSONResult
		Invest() helper.JSONResult
		Disburse() helper.JSONResult
	}
)

func NewUsecase(opts Usecase) IUsecase {
	return &Usecase{
		LoanRepository:     opts.LoanRepository,
		DocumentRepository: opts.DocumentRepository,
		Logger:             opts.Logger,
	}
}

package loan

import (
	"github.com/Gash21/amartha-test/internal/application/document"
	"github.com/Gash21/amartha-test/internal/application/loan"
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
		Approved() helper.JSONResult
		Invested() helper.JSONResult
		Disbursed() helper.JSONResult
	}
)

func NewUsecase(opts Usecase) IUsecase {
	return &Usecase{
		LoanRepository:     opts.LoanRepository,
		DocumentRepository: opts.DocumentRepository,
		Logger:             opts.Logger,
	}
}

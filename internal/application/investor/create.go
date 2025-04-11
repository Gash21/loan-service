package investor

import (
	"context"
	"net/http"

	"github.com/Gash21/amartha-test/internal/application/investor/dto"
	"github.com/Gash21/amartha-test/internal/domain/investor"
	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) Create(ctx context.Context, req dto.CreateRequest) helper.JSONResult {
	investor := u.InvestorRepository.Create(&investor.Investor{
		Name: req.Name,
	})

	return helper.ResponseSuccess(http.StatusCreated, investor)
}

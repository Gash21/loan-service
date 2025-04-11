package borrower

import (
	"context"
	"net/http"

	"github.com/Gash21/amartha-test/internal/application/borrower/dto"
	"github.com/Gash21/amartha-test/internal/domain/borrower"
	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) Create(ctx context.Context, req dto.CreateRequest) helper.JSONResult {
	borrower := u.BorrowerRepository.Create(&borrower.Borrower{
		Name: req.Name,
	})

	return helper.ResponseSuccess(http.StatusCreated, borrower)
}

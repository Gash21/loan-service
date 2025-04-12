package loan

import (
	"context"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) List(ctx context.Context, req *dto.ListRequest) helper.JSONResult {
	loans, total := u.LoanRepository.FindPaginated(req.Page, req.Limit, req.Status)
	listResponse := &dto.ListResponse{}
	listResponse = listResponse.FromModel(&u.Config, &loans)
	return helper.ResponsePagination(req.Page, req.Limit, len(loans), int(total), listResponse.Data)
}

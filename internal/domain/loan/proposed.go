package loan

import (
	"net/http"

	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) Propose() helper.JSONResult {
	return helper.ResponseSuccess(http.StatusOK, "Loan proposed successfully")
}

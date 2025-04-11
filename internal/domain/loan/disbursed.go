package loan

import (
	"net/http"

	"github.com/Gash21/amartha-test/internal/shared/helper"
)

func (u *Usecase) Disbursed() helper.JSONResult {
	return helper.ResponseSuccess(http.StatusOK, "Loan disbursed successfully")
}

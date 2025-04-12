package loan

import (
	"net/http"
	"time"

	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Approve(c *fiber.Ctx) error {
	m := &dto.ApproveRequest{}
	helper.BindModel(h.Logger, c, m)

	// bind files
	form, _ := c.MultipartForm()
	files := form.File["evidence"]
	DateOfApproval, err := time.Parse("2006-01-02", m.DateString)
	if err != nil {
		perr := helper.ResponseFailed(http.StatusBadRequest, "Invalid date format", nil)
		return c.Status(perr.Code).JSON(perr)
	}
	m.DateOfApproval = DateOfApproval

	// check if file is uploaded
	if len(files) == 0 {
		perr := helper.ResponseFailed(http.StatusBadRequest, "No file uploaded", nil)
		return c.Status(perr.Code).JSON(perr)
	}

	m.Evidence = files[0]

	uc := h.Usecase.Approve(c.UserContext(), m)
	return c.Status(uc.Code).JSON(uc)
}

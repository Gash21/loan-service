package loan

import (
	"github.com/Gash21/amartha-test/internal/application/loan/dto"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Invest(c *fiber.Ctx) error {
	req := &dto.InvestRequest{}

	if err := helper.BindAndValidateModel(h.Logger, c, h.Validator, req); err != nil {
		return c.Status(err.Code).JSON(err.ResponseBody)
	}
	uc := h.Usecase.Invest(c.UserContext(), req)
	return c.Status(uc.Code).JSON(uc)
}

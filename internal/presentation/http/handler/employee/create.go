package employee

import (
	"github.com/Gash21/amartha-test/internal/application/employee/dto"
	"github.com/Gash21/amartha-test/internal/shared/helper"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Create(c *fiber.Ctx) error {
	m := dto.CreateRequest{}

	if err := helper.BindAndValidateModel(h.Logger, c, h.Validator, &m); err != nil {
		return c.Status(err.Code).JSON(err.ResponseBody)
	}

	uc := h.Usecase.Create(c.UserContext(), m)
	return c.Status(uc.Code).JSON(uc)
}

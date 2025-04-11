package loan

import "github.com/gofiber/fiber/v2"

func (h *Handler) Propose(c *fiber.Ctx) error {
	return c.JSON("Halo")
}

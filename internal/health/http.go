package health

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterHandler(srv *fiber.App) {
	srv.Get("/health", health)
}

// HealthCheck godoc
// @Summary      Determines if a service is operating
// @Success      200
// @Router       /health [get]
func health(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

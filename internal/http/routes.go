package http

import "github.com/gofiber/fiber/v2"

func (s *Server) WithPartnerApiRoutes() {
	s.FiberApp.Post("/balance", s.mockPartnerController.balance)
	s.FiberApp.Post("/bet", s.mockPartnerController.bet)
	s.FiberApp.Post("/win", s.mockPartnerController.win)
	s.FiberApp.Post("/rollback", s.mockPartnerController.notImplementedYet)
}

func (s *Server) WithMockRoutes() {
	s.FiberApp.Get("/start", s.interactionController.start)
	s.FiberApp.Get("/reset", s.interactionController.reset)
	s.FiberApp.Get("/demo-game", s.interactionController.demoGame)
}

func (s *Server) WithHealth() {
	s.FiberApp.Get("/health", health)
}

// HealthCheck godoc
// @Summary      Determines if a service is operating
// @Success      200
// @Router       /health [get]
func health(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

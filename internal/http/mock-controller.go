package http

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/partner"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type MockController struct {
	partnerService *partner.Service
}

func NewMockController(p *partner.Service) *MockController {
	return &MockController{partnerService: p}
}

func (s *MockController) start(c *fiber.Ctx) error {
	game, err := s.partnerService.PostLaunchGame()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	body, _ := json.Marshal(game)
	return c.Send(body)
}

func (s *MockController) reset(c *fiber.Ctx) error {
	s.partnerService.Reset()
	return c.SendStatus(fiber.StatusOK)
}

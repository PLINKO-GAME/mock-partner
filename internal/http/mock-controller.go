package http

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/partner"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

type MockController struct {
	partnerService *partner.Service
}

func NewMockController(p *partner.Service) *MockController {
	return &MockController{partnerService: p}
}

func (s *MockController) start(c *fiber.Ctx) error {
	game, err := s.startGame()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(game)
}

func (s *MockController) demoGame(c *fiber.Ctx) error {
	game, err := s.startGame()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	split := strings.SplitAfter(game.URL, "?token=")
	if len(split) != 2 {
		log.Errorf("received bad URL from core: [%s]", game.URL)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(&dto.DemoLaunchGameResponse{OneTimeToken: split[1]})
}

func (s *MockController) startGame() (*dto.LaunchGameResponse, error) {
	game, err := s.partnerService.PostLaunchGame()
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s *MockController) reset(c *fiber.Ctx) error {
	s.partnerService.Reset()
	return c.SendStatus(fiber.StatusOK)
}

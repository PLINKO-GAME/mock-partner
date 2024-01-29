package http

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/partner"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type PartnerApiController struct {
	partnerService *partner.Service
}

func NewPartnerApiController(p *partner.Service) *PartnerApiController {
	return &PartnerApiController{partnerService: p}
}

func (s *PartnerApiController) balance(c *fiber.Ctx) error {
	payload := dto.BalanceRequest{}
	if err := c.BodyParser(&payload); err != nil {
		log.WithError(err).Warn("failed to get balance")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	balance, err := s.partnerService.GetBalance(&payload)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(balance)
}

func (s *PartnerApiController) bet(c *fiber.Ctx) error {
	payload := dto.BetRequest{}
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	balance, err := s.partnerService.Bet(&payload)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(balance)
}

func (s *PartnerApiController) win(c *fiber.Ctx) error {
	payload := dto.WinRequest{}
	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	balance, err := s.partnerService.Win(&payload)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(balance)
}

func (s *PartnerApiController) notImplementedYet(c *fiber.Ctx) error {
	panic("not implemented yet")
}

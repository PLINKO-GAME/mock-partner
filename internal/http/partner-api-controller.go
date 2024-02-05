package http

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/partner"
	"bitbucket.org/1-pixel-games/mock-partner/internal/sign"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type PartnerApiController struct {
	partnerService *partner.Service
	signService    *sign.Service
}

func NewPartnerApiController(p *partner.Service, s *sign.Service) *PartnerApiController {
	return &PartnerApiController{partnerService: p, signService: s}
}

func (s *PartnerApiController) balance(c *fiber.Ctx) error {
	if !s.signService.VerifyProviderSignature(c) {
		return c.SendStatus(fiber.StatusForbidden)
	}

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
	if !s.signService.VerifyProviderSignature(c) {
		log.Error("bad signature provided for /bet request")
		return c.SendStatus(fiber.StatusForbidden)
	}

	payload := dto.BetRequest{}
	if err := c.BodyParser(&payload); err != nil {
		log.Error("failed to parse /bet request body")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	balance, err := s.partnerService.Bet(&payload)
	if err != nil {
		log.WithError(err).Error("failed to handle bet")
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(balance)
}

func (s *PartnerApiController) win(c *fiber.Ctx) error {
	if !s.signService.VerifyProviderSignature(c) {
		return c.SendStatus(fiber.StatusForbidden)
	}

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

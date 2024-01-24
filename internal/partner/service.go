package partner

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/session"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Service struct {
	coreURL        string
	sessionService *session.Service
}

func New(coreURL string) *Service {
	return &Service{
		coreURL:        coreURL,
		sessionService: session.New(),
	}
}

func (s *Service) RegisterHandler(srv *fiber.App) {
	srv.Get("/start", s.start)
	srv.Post("/balance", s.balance)
	srv.Post("/bet", s.start)
	srv.Post("/win", s.start)
	srv.Post("/rollback", s.start)
}

func (s *Service) start(c *fiber.Ctx) error {
	return c.SendStatus(s.postLaunchGame())
}

func (s *Service) postLaunchGame() int {
	se := s.sessionService.GenerateAndStoreSession()
	data := dto.LaunchGameRequest{
		Currency:     se.Currency,
		SessionToken: se.Token,
		PlayerId:     se.PlayerID,
		GameId:       "1",
		Language:     "en",
	}
	body, _ := json.Marshal(data)

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/api/launch-game", s.coreURL), bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return fiber.StatusBadRequest
	}

	fmt.Printf("Game launched. player: [%s] session: [%s] status: [%d]\n", data.PlayerId, data.SessionToken, res.StatusCode)

	return res.StatusCode
}

func (s *Service) balance(c *fiber.Ctx) error {
	payload := dto.BalanceRequest{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	balance, err := s.sessionService.GetBalance(payload.Token)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(&dto.BalanceResponse{
		Balance:   balance,
		Timestamp: 0,
	})
}

func (s *Service) bet(c *fiber.Ctx) error {
	payload := dto.BetRequest{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	balance, err := s.sessionService.Bet(payload.Token, payload.BetAmount)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(&dto.BalanceResponse{
		Balance:   balance,
		Timestamp: 0,
	})
}

func (s *Service) win(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (s *Service) rollback(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

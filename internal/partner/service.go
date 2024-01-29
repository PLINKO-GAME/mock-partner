package partner

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/session"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
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

func (s *Service) PostLaunchGame() int {
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

	log.Infof("Game launched. player: [%s] session: [%s] status: [%d]\n", data.PlayerId, data.SessionToken, res.StatusCode)

	return res.StatusCode
}

func (s *Service) GetBalance(request *dto.BalanceRequest) (*dto.BalanceResponse, error) {
	balance, err := s.sessionService.GetBalance(request.Token)
	if err != nil {
		return nil, errors.New("token not found")
	}

	return &dto.BalanceResponse{
		Balance:   balance,
		Timestamp: time.Now().UnixMicro(),
	}, nil
}

func (s *Service) Bet(request *dto.BetRequest) (*dto.BalanceResponse, error) {
	balance, err := s.sessionService.Bet(request.Token, request.BetAmount)
	if err != nil {
		return nil, err
	}

	return &dto.BalanceResponse{
		Balance:   balance,
		Timestamp: time.Now().UnixMicro(),
	}, nil
}

func (s *Service) Win(request *dto.WinRequest) (*dto.BalanceResponse, error) {
	balance, err := s.sessionService.Win(request.Token, request.WinAmount)
	if err != nil {
		return nil, err
	}

	return &dto.BalanceResponse{
		Balance:   balance,
		Timestamp: time.Now().UnixMicro(),
	}, nil
}

func (s *Service) rollback(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func (s *Service) Reset() {
	s.sessionService.Reset()
}

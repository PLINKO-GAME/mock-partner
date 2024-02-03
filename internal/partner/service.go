package partner

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/session"
	"bitbucket.org/1-pixel-games/mock-partner/internal/sign"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type Service struct {
	coreURL        string
	sessionService *session.Service
	signService    *sign.Service
}

func New(signService *sign.Service, coreURL string) *Service {
	return &Service{
		coreURL:        coreURL,
		sessionService: session.New(),
		signService:    signService,
	}
}

func (s *Service) PostLaunchGame() (*dto.LaunchGameResponse, error) {
	se := s.sessionService.GenerateAndStoreSession()
	data := dto.LaunchGameRequest{
		OperatorId:   sign.OperatorID,
		Currency:     se.Currency,
		SessionToken: se.Token,
		PlayerId:     se.PlayerID,
		GameId:       "1",
		Language:     "en",
	}
	body, _ := json.Marshal(data)

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/launch-game", s.coreURL), bytes.NewBuffer(body))
	if err != nil {
		log.WithError(err).Fatal("failed to form http request")
	}

	r.Header.Add("Content-Type", "application/json")
	s.signService.AttachOperatorSignature(r, body)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		log.WithError(err).Fatal("/launch-game request failed")
		return nil, err
	}
	// TODO check res.statusCode too

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.WithError(err).Fatal("/launch-game failed to read response")
		return nil, err
	}

	var clgr dto.CoreLaunchGameResponse
	err = json.Unmarshal(resBody, &clgr)

	log.Infof("Game launched. player: [%s] session: [%s] status: [%d] url: [%s]\n", data.PlayerId, data.SessionToken, res.StatusCode, clgr.URL)

	return &dto.LaunchGameResponse{
		URL:      clgr.URL,
		PlayerID: se.PlayerID,
	}, nil
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
		log.WithError(err).Fatal("payout failed")
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

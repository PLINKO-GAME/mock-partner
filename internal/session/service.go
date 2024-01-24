package session

import (
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/shopspring/decimal"
	"strconv"
)

type Service struct {
	data map[string]*Session
}

type Session struct {
	Token    string
	PlayerID string
	Balance  decimal.Decimal
	Currency string
}

func New() *Service {
	return &Service{
		data: make(map[string]*Session),
	}
}

func (s *Service) GenerateAndStoreSession() *Session {
	session := &Session{
		Token:    gofakeit.UUID(),
		PlayerID: strconv.Itoa(gofakeit.Number(1, 1000000)),
		Balance:  decimal.NewFromInt(900_000),
		Currency: "EUR",
	}

	s.data[session.Token] = session

	return session
}

func (s *Service) GetBalance(token string) (decimal.Decimal, error) {
	val, ok := s.data[token]
	if !ok {
		return decimal.Zero, errors.New(fmt.Sprintf("token not found %s", token))
	}

	return val.Balance, nil
}

func (s *Service) Bet(token string, betAmount decimal.Decimal) (decimal.Decimal, error) {
	val, ok := s.data[token]
	if !ok {
		return decimal.Zero, errors.New(fmt.Sprintf("token not found %s", token))
	}

	if val.Balance.GreaterThanOrEqual(betAmount) {
		val.Balance.Sub(betAmount)
	} else {
		return val.Balance, errors.New("not enough funds")
	}

	return val.Balance, nil
}

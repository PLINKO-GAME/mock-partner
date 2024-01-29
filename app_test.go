package main

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/session"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jarcoal/httpmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http/httptest"
	"testing"
)

var conf = config{CoreURL: "http://localhost:1234"}

func Test_health(t *testing.T) {
	a, err := newApplication(&config{})
	assert.Nil(t, err)

	req := httptest.NewRequest("GET", "/health", nil)
	test, err := a.server.FiberApp.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, test.StatusCode)
}

func Test_start(t *testing.T) {
	a, err := newApplication(&conf)
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/launch-game", conf.CoreURL),
		httpmock.NewStringResponder(200, ""))

	req := httptest.NewRequest("GET", "/start", nil)
	test, err := a.server.FiberApp.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, test.StatusCode)
}

func Test_bet(t *testing.T) {
	a, err := newApplication(&conf)
	assert.Nil(t, err)

	body, err := json.Marshal(&dto.BetRequest{
		Token:     session.TemporaryToken,
		BetAmount: decimal.NewFromInt(1000),
	})
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/bet", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	test, err := a.server.FiberApp.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, test.StatusCode)

	read, err := io.ReadAll(test.Body)
	defer test.Body.Close()
	assert.Nil(t, err)
	var balanceResponse dto.BalanceResponse
	err = json.Unmarshal(read, &balanceResponse)

	assert.Equal(t, decimal.NewFromInt(999_999_000), balanceResponse.Balance)
}

func Test_win(t *testing.T) {
	a, err := newApplication(&conf)
	assert.Nil(t, err)

	body, err := json.Marshal(&dto.WinRequest{
		Token:     session.TemporaryToken,
		WinAmount: decimal.NewFromInt(700),
	})
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/win", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	test, err := a.server.FiberApp.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, test.StatusCode)

	read, err := io.ReadAll(test.Body)
	defer test.Body.Close()
	assert.Nil(t, err)
	var balanceResponse dto.BalanceResponse
	err = json.Unmarshal(read, &balanceResponse)

	assert.Equal(t, decimal.NewFromInt(1_000_000_700), balanceResponse.Balance)
}

package main

import (
	"bitbucket.org/1-pixel-games/mock-partner/internal/dto"
	"bitbucket.org/1-pixel-games/mock-partner/internal/session"
	"bitbucket.org/1-pixel-games/mock-partner/internal/sign"
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

func getTestConfig() *config {
	conf, err := newConfig()
	if err != nil {
		panic(err)
	}
	conf.CoreURL = "http://localhost:1234"
	conf.PublicKey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhcnimjhkkGue02/EdYkV\n47tR0L36Wcunn/1WW4noRLD8M04Y9G36UT2uzt20u5KzKwOk1NdW0/0WYqQCMJ4D\nfKVdaLiXpZFsLsirT8r8cC0mpuNAAUHdl/bGDJa9kFrbwT6RyTkIMAxLnG16yeiu\nam/799ve9V04kDvC2F/YFCX1SumrJCVgbyw9K/xpYWp7Fppt/pB3KXlvZf0FyIDk\nI9j767tMfuTlOVLPY/n5x+ZaVmF7vxhKJknoboHY+9v5KaU7XeTrThZfSl0LjboB\nBpFl1gWpKNglGQbOvjOdijV51xgi8IeIZ90u2W3LtTbL+k+c16qMVOaFA3sJVQwZ\nCwIDAQAB\n-----END PUBLIC KEY-----"
	conf.PrivateKey = "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAhcnimjhkkGue02/EdYkV47tR0L36Wcunn/1WW4noRLD8M04Y\n9G36UT2uzt20u5KzKwOk1NdW0/0WYqQCMJ4DfKVdaLiXpZFsLsirT8r8cC0mpuNA\nAUHdl/bGDJa9kFrbwT6RyTkIMAxLnG16yeiuam/799ve9V04kDvC2F/YFCX1Sumr\nJCVgbyw9K/xpYWp7Fppt/pB3KXlvZf0FyIDkI9j767tMfuTlOVLPY/n5x+ZaVmF7\nvxhKJknoboHY+9v5KaU7XeTrThZfSl0LjboBBpFl1gWpKNglGQbOvjOdijV51xgi\n8IeIZ90u2W3LtTbL+k+c16qMVOaFA3sJVQwZCwIDAQABAoIBAAJ/JPppyP6k837Q\nnCLxXvYz/a/ei7h3Q3aJ3L2ykiIOB3bRo0eUcdJoS0XS/1dswmkwFThfmGA2Xd+T\nXfMYT8pYr6iPoUzWrOUmm4POru1M+mas4PnlB8SZN1Lu0TTLbURq7X+Kz+tNn2+Y\n32y7Kd4UnugeM0fy6GZQpy8wgrDFiWKphSB7pz/czovAR7BQLbCIhpGgQc4LF22D\n1EAx5i/wSU5b8KgGeiOsdpT9LzOaFosaeq1tDGlwpN4JKuj6sSgvhST96f/XihD5\nnwBNIDiM5IFMr4MplJThr2oFTQH7TyK6/bswvvURReyXiKpVaaZ/QTyY2hQ3yrjV\nFZGPX2ECgYEA9StqKNYmqg8W+X6MUE0EyAlHUHey30dgz1nn3uHmynhQMWEWc2X0\ni7r17h1Yi37DN0veIJRW3e1/DDypezIebQkaoyGmuy2Ogh3WCrH/nL/WVMgcAIXu\nLRS111M4LtJFOc5vbCw8U/34OFH0T6k5/8gjtrcx12ExZzezoSfXiBECgYEAi7Lh\nTmRwOHcYvMH84RETb4BoBq3V97CCO91ZpinDQa8z9v3tsEl5NhY0N9cceMvnB/+f\nqQ3tNHCrZcktXLXwLLPDCtYU9BpY/sRZy39kCsEo8t8xtGPjiI+rH0AhheTgUTCc\n5TK3LwKmlkUGGDpww9eOIFljbIMeoH2+HOU2C1sCgYA8SfTNHfxcDWHk8I2ooYfv\nePikfQrrhS31T3KJiJusZnGx8uIGdqfwRIV9jJHdm8p9qpZxBIloAaMgazpyJRz+\nSyLVwsyxcr58mMGt15+3+CTIrHzWVBkB1PnyfXBvcx263VzhCO+859NGZkDh5gdx\nMtI1eE81W50+eKAfnSCPQQKBgE7oWW9IOEMMsoJcKJSQaqP+qcOsCUIBB279FphO\n2qWNaxLGV63NspOkcxZfgQuSUQspjmuVHDkUsxupSOAnPGRjnXXPesJu53nwOrBB\nYqbYeGLHQ3IbQfhu/j+Gn+jbYQE7LkQgI2yAWMxkbI7e47cbWIJZO1mdrn0EyY/U\nwHQlAoGAMqsSsgQM4IKzkK3kW+WmCEtpfPu+A6gaDAPVWdY1UU/r5IfGSV47r+bE\nj2BQaPQFiLzqA4SRDuvZHDRyI3MsqT9t1Mx6fBJ+x2kb3Muctc/SR2JpsIgbB1xv\nKk1vBRzeosnDbvbnZVkFih/lWLSEuLyE2iqy+kwScwIhiBcN/a8=\n-----END RSA PRIVATE KEY-----"

	return conf
}

func Test_health(t *testing.T) {
	a, err := newApplication(&config{})
	assert.Nil(t, err)

	req := httptest.NewRequest("GET", "/health", nil)
	test, err := a.server.FiberApp.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, test.StatusCode)
}

func Test_start(t *testing.T) {
	conf := getTestConfig()
	a, err := newApplication(conf)
	assert.Nil(t, err)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", fmt.Sprintf("%s/api/v1/launch-game", conf.CoreURL),
		httpmock.NewStringResponder(200, ""))

	req := httptest.NewRequest("GET", "/start", nil)
	test, err := a.server.FiberApp.Test(req)

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, test.StatusCode)
}

func Test_bet_withSignature(t *testing.T) {
	conf := getTestConfig()
	a, err := newApplication(conf)
	assert.Nil(t, err)

	body, err := json.Marshal(&dto.BetRequest{
		Token:     session.TemporaryToken,
		BetAmount: decimal.NewFromInt(1000),
	})
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/bet", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	signService := sign.New(conf.PrivateKey, conf.PublicKey)
	signService.AttachOperatorSignature(req, body)

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

func Test_bet_withoutSignature(t *testing.T) {
	conf := getTestConfig()
	a, err := newApplication(conf)
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
	assert.Equal(t, fiber.StatusForbidden, test.StatusCode)
}

func Test_win(t *testing.T) {
	conf := getTestConfig()
	a, err := newApplication(conf)
	assert.Nil(t, err)

	body, err := json.Marshal(&dto.WinRequest{
		Token:     session.TemporaryToken,
		WinAmount: decimal.NewFromInt(700),
	})
	assert.Nil(t, err)

	req := httptest.NewRequest("POST", "/win", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	signService := sign.New(conf.PrivateKey, conf.PublicKey)
	signService.AttachOperatorSignature(req, body)
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

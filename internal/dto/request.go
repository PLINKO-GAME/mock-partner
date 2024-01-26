package dto

import "github.com/shopspring/decimal"

type LaunchGameRequest struct {
	Currency               string `json:"currency"`
	SessionToken           string `json:"session_token"`
	PlayerId               string `json:"player_id"`
	GameId                 string `json:"game_id"`
	CashierUrl             string `json:"cashier_url"`
	LobbyUrl               string `json:"lobby_url"`
	ResponsibleGamblingUrl string `json:"responsible_gambling_url"`
	ExitUrl                string `json:"exit_url"`
	Language               string `json:"language"`
	UserIP                 string `json:"user_ip"`
	UserCountry            string `json:"user_country"`
}

type BalanceRequest struct {
	Token string `json:"token"`
}

type BetRequest struct {
	Token     string          `json:"token"`
	BetAmount decimal.Decimal `json:"bet_amount"`
}

type WinRequest struct {
	Token     string          `json:"token"`
	WinAmount decimal.Decimal `json:"win_amount"`
}

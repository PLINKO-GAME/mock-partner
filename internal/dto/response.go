package dto

import (
	"github.com/shopspring/decimal"
)

type BalanceResponse struct {
	Balance   decimal.Decimal
	Timestamp int64 // microseconds
}

type LaunchGameResponse struct {
	URL      string `json:"url"`
	PlayerID string `json:"player_id"`
}

type CoreLaunchGameResponse struct {
	URL string `json:"url"`
}

type DemoLaunchGameResponse struct {
	OneTimeToken string `json:"token"`
}

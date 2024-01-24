package dto

import (
	"github.com/shopspring/decimal"
)

type BalanceResponse struct {
	Balance   decimal.Decimal
	Timestamp int64 // microseconds
}

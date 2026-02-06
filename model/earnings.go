package model

import (
	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetEarningsCalendarParams struct {
	Since *types.Date `query:"from"`
	Until *types.Date `query:"to"`
}

type GetEarningsCalendarResponse struct {
	Date             types.Date       `json:"date"`
	Symbol           string           `json:"symbol"`
	EPS              *decimal.Decimal `json:"eps"`
	EPSEstimated     *decimal.Decimal `json:"epsEstimated"`
	Revenue          *decimal.Decimal `json:"revenue"`
	RevenueEstimated *decimal.Decimal `json:"revenueEstimated"`
	LastUpdatedAt    types.Date       `json:"lastUpdated"`
}

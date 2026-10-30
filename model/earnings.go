package model

import (
	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetEarningsCalendarParams struct {
	Since *types.Date `query:"from"`
	Until *types.Date `query:"to"`
}

// GetEarningsCalendarResponse is one /stable/earnings-calendar row. EPS and
// Revenue are the reported actuals, nil until the company reports; the stable
// API names them epsActual and revenueActual (the v3 API's eps and revenue
// are never sent).
type GetEarningsCalendarResponse struct {
	Date             types.Date       `json:"date"`
	Symbol           string           `json:"symbol"`
	EPS              *decimal.Decimal `json:"epsActual"`
	EPSEstimated     *decimal.Decimal `json:"epsEstimated"`
	Revenue          *decimal.Decimal `json:"revenueActual"`
	RevenueEstimated *decimal.Decimal `json:"revenueEstimated"`
	LastUpdatedAt    types.Date       `json:"lastUpdated"`
}

package model

import (
	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetHistoricalEarningsCalendarParams struct {
	Symbol string      `path:"symbol"`
	Since  *types.Date `query:"since"`
	Until  *types.Date `query:"until"`
}

type GetEarningsCalendarParams struct {
	Since *types.Date `query:"since"`
	Until *types.Date `query:"until"`
}

type GetEarningsCalendarResponse struct {
	Date             types.Date       `json:"date"`
	Symbol           string           `json:"symbol"`
	EPS              *decimal.Decimal `json:"eps"`
	EPSEstimated     *decimal.Decimal `json:"epsEstimated"`
	MarketTime       string           `json:"time"`
	Revenue          *decimal.Decimal `json:"revenue"`
	RevenueEstimated *decimal.Decimal `json:"revenueEstimated"`
	FiscalDateEnding types.Date       `json:"fiscalDateEnding"`
	UpdatedFromDate  types.Date       `json:"updatedFromDate"`
}

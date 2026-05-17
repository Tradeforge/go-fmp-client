package model

import (
	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetDividendsCalendarParams struct {
	Since *types.Date `query:"from"`
	Until *types.Date `query:"to"`
}

type GetDividendsCalendarResponse struct {
	Symbol          string                    `json:"symbol"`
	Date            types.Date                `json:"date"`
	RecordDate      types.EmptyOr[types.Date] `json:"recordDate"`
	PaymentDate     types.EmptyOr[types.Date] `json:"paymentDate"`
	DeclarationDate types.EmptyOr[types.Date] `json:"declarationDate"`
	AdjDividend     decimal.Decimal           `json:"adjDividend"`
	Dividend        decimal.Decimal           `json:"dividend"`
	Yield           decimal.Decimal           `json:"yield"`
	Frequency       string                    `json:"frequency"`
}

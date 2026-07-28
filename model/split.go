package model

import (
	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetSplitsCalendarParams struct {
	Since *types.Date `query:"from"`
	Until *types.Date `query:"to"`
}

// GetSplitsCalendarResponse is a single row of FMP's splits calendar.
//
// The split ratio is expressed as Numerator:Denominator — a 2:1 forward split
// has Numerator=2, Denominator=1, while a 1:5 reverse split has Numerator=1,
// Denominator=5. Upstream returns whole numbers in practice, but they are
// modelled as decimals to match the rest of the FMP numeric surface.
type GetSplitsCalendarResponse struct {
	Symbol      string          `json:"symbol"`
	Date        types.Date      `json:"date"`
	Numerator   decimal.Decimal `json:"numerator"`
	Denominator decimal.Decimal `json:"denominator"`
	// SplitType distinguishes a true split from a stock dividend.
	// Observed values: "stock-split", "stock-dividend". Empty when upstream omits it.
	SplitType string `json:"splitType"`
}

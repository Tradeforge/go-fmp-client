package model

import (
	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetHistoricalBarsParams struct {
	Timeframe Timeframe  `path:"timeframe,required"`
	Symbol    string     `query:"symbol,required"`
	Since     types.Date `query:"from,omitempty"`
	Until     types.Date `query:"to,omitempty"`
}

type GetHistoricalBarsResponse = []Bar

type Bar struct {
	Open     decimal.Decimal `json:"open"`
	High     decimal.Decimal `json:"high"`
	Low      decimal.Decimal `json:"low"`
	Close    decimal.Decimal `json:"close"`
	Volume   decimal.Decimal `json:"volume"`
	DateTime types.DateTime  `json:"date"`
}

type GetHistoricalPricesEODParams struct {
	Symbol string     `query:"symbol,required"`
	Since  types.Date `query:"from,omitempty"`
	Until  types.Date `query:"to,omitempty"`
}

type GetHistoricalPricesEODResponse = []HistoricalPriceEOD

type HistoricalPriceEOD struct {
	Symbol        string          `json:"symbol"`
	Date          types.Date      `json:"date"`
	Open          decimal.Decimal `json:"open"`
	High          decimal.Decimal `json:"high"`
	Low           decimal.Decimal `json:"low"`
	Close         decimal.Decimal `json:"close"`
	Volume        decimal.Decimal `json:"volume"`
	Change        decimal.Decimal `json:"change"`
	ChangePercent decimal.Decimal `json:"changePercent"`
	VWAP          decimal.Decimal `json:"vwap"`
}

type GetQuoteParams struct {
	Symbol string `query:"symbol,required"`
}

type GetQuoteResponse = TickerPrice

type BatchGetQuoteParams struct {
	Symbols string `query:"symbols,required"`
}

type BatchGetQuoteResponse = []TickerPrice

type BatchGetQuotesByExchangeParams struct {
	Exchange string `query:"exchange,required"`
	Short    bool   `query:"short"` // DO NOT USE: we always expect full prices to be returned.
}

type BatchGetQuotesByExchangeResponse = []TickerPrice

type TickerPrice struct {
	Symbol           string          `json:"symbol"`
	Name             string          `json:"name"`
	Open             decimal.Decimal `json:"open"`
	Price            decimal.Decimal `json:"price"`
	PreviousClose    decimal.Decimal `json:"previousClose"`
	ChangePercentage decimal.Decimal `json:"changePercentage"`
	Change           decimal.Decimal `json:"change"`
	DayLow           decimal.Decimal `json:"dayLow"`
	DayHigh          decimal.Decimal `json:"dayHigh"`
	YearLow          decimal.Decimal `json:"yearLow"`
	YearHigh         decimal.Decimal `json:"yearHigh"`
	PriceAvg50       decimal.Decimal `json:"priceAvg50"`
	PriceAvg200      decimal.Decimal `json:"priceAvg200"`
	MarketCap        decimal.Decimal `json:"marketCap"`
	Exchange         string          `json:"exchange"`
	Volume           decimal.Decimal `json:"volume"`
	Timestamp        int64           `json:"timestamp"`
}

type GetPriceChangeParams struct {
	Symbol string `query:"symbol,required"`
}

type GetPriceChangeResponse struct {
	Symbol    string          `json:"symbol"`
	Change1D  decimal.Decimal `json:"1D"`
	Change5D  decimal.Decimal `json:"5D"`
	Change1M  decimal.Decimal `json:"1M"`
	Change3M  decimal.Decimal `json:"3M"`
	Change6M  decimal.Decimal `json:"6M"`
	Change1Y  decimal.Decimal `json:"1Y"`
	Change3Y  decimal.Decimal `json:"2Y"`
	Change5Y  decimal.Decimal `json:"5Y"`
	Change10Y decimal.Decimal `json:"10Y"`
	ChangeYTD decimal.Decimal `json:"ytd"`
	ChangeMax decimal.Decimal `json:"max"`
}

type GetHistoricalMarketCapParams struct {
	Symbol string     `query:"symbol,required"`
	Since  types.Date `query:"from,omitempty"`
	Until  types.Date `query:"to,omitempty"`
}

type GetHistoricalMarketCapResponse []HistoricalMarketCap

type HistoricalMarketCap struct {
	Symbol string          `json:"symbol"`
	Date   types.Date      `json:"date"`
	Value  decimal.Decimal `json:"marketCap"`
}

type GetBulkPriceEODParams struct {
	Date types.Date `query:"date,required"`
}

type GetBulkPriceEODResponse = []BulkPriceEOD

type BulkPriceEOD struct {
	Symbol   string          `json:"symbol"`
	Date     types.Date      `json:"date"`
	Open     decimal.Decimal `json:"open"`
	Low      decimal.Decimal `json:"low"`
	High     decimal.Decimal `json:"high"`
	Close    decimal.Decimal `json:"close"`
	AdjClose decimal.Decimal `json:"adjClose"`
	Volume   decimal.Decimal `json:"volume"`
}

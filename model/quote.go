package model

import "github.com/shopspring/decimal"

type GetRealTimeQuoteParams struct {
	Symbol string `path:"symbol,required"`
}

type GetRealTimeQuoteResponse = TickerQuote

type BatchGetRealTimeQuoteParams struct {
	Symbols string `path:"symbols,required"`
}

type BatchGetRealTimeQuoteResponse = []TickerQuote

type TickerQuote struct {
	Symbol      string          `json:"symbol"`
	BidPrice    decimal.Decimal `json:"bidPrice"`
	AskPrice    decimal.Decimal `json:"askPrice"`
	BidSize     uint64          `json:"bidSize"`
	AskSize     uint64          `json:"askSize"`
	Volume      decimal.Decimal `json:"volume"`
	LastUpdated int64           `json:"lastUpdated"`
}

type GetFullPriceParams struct {
	Symbol string `path:"symbol,required"`
}

type GetFullPriceResponse = TickerPrice

type BatchGetFullPriceParams struct {
	Symbols string `path:"symbols,required"`
}

type BatchGetFullPriceResponse = []TickerPrice

type TickerPrice struct {
	Symbol           string          `json:"symbol"`
	Name             string          `json:"name"`
	Open             decimal.Decimal `json:"open"`
	Price            decimal.Decimal `json:"price"`
	PreviousClose    decimal.Decimal `json:"previousClose"`
	ChangePercentage decimal.Decimal `json:"changesPercentage"`
	Change           decimal.Decimal `json:"change"`
	DayLow           decimal.Decimal `json:"dayLow"`
	DayHigh          decimal.Decimal `json:"dayHigh"`
	YearLow          decimal.Decimal `json:"yearLow"`
	YearHigh         decimal.Decimal `json:"yearHigh"`
	PriceAvg50       decimal.Decimal `json:"priceAvg50"`
	PriceAvg200      decimal.Decimal `json:"priceAvg200"`
	Volume           decimal.Decimal `json:"volume"`
	AvgVolume        decimal.Decimal `json:"avgVolume"`
	MarketCap        decimal.Decimal `json:"marketCap"`
	Exchange         string          `json:"exchange"`
	Timestamp        int64           `json:"timestamp"`
}

type BatchGetPriceChangeParams struct {
	Symbols string `path:"symbols,required"`
}

type GetPriceChangeParams struct {
	Symbol string `path:"symbol,required"`
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

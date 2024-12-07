package model

import "github.com/shopspring/decimal"

type GetRealtimeQuoteParams struct {
	Symbol string `path:"symbol,required"`
}

type GetRealtimeQuoteResponse = TickerQuote

type BatchGetRealtimeQuoteParams struct {
	Symbols string `path:"symbols,required"`
}

type BatchGetRealtimeQuoteResponse = []TickerQuote

type TickerQuote struct {
	Symbol      string          `json:"symbol"`
	BidPrice    decimal.Decimal `json:"bidPrice"`
	AskPrice    decimal.Decimal `json:"askPrice"`
	BidSize     decimal.Decimal `json:"bidSize"`
	AskSize     decimal.Decimal `json:"askSize"`
	Volume      decimal.Decimal `json:"volume"`
	LastUpdated int64           `json:"lastUpdated"`
}

type ListAllRealtimeQuotesResponse = []TickerQuote

type ListExchangeSymbolsParams struct {
	Exchange string `path:"exchange,required"`
}

type ListExchangeSymbolsResponse = []TickerPrice

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
	MarketCap        decimal.Decimal `json:"marketCap"`
	PE               decimal.Decimal `json:"pe"`
	EPS              decimal.Decimal `json:"eps"`
	Exchange         string          `json:"exchange"`
	Volume           decimal.Decimal `json:"volume"`
	AvgVolume        decimal.Decimal `json:"avgVolume"`
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

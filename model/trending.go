package model

import "github.com/shopspring/decimal"

type ListMostActiveTickersResponse = []PartialTicker

type PartialTicker struct {
	Symbol        string          `json:"symbol"`
	Name          string          `json:"name"`
	Price         decimal.Decimal `json:"price"`
	Change        decimal.Decimal `json:"change"`
	ChangePercent decimal.Decimal `json:"changePercent"`
}

type ListGainersResponse = []PartialTicker

type ListLosersResponse = []PartialTicker

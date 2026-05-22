package model

import (
	"go.tradeforge.dev/fmp/pkg/types"
)

type GetIndexListResponse = []IndexListEntry

type IndexListEntry struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exchange string `json:"exchange"`
	Currency string `json:"currency"`
}

type BatchGetIndexQuotesResponse = []TickerQuote

type BatchGetIndexShortQuotesResponse = []TickerShortQuote

type GetHistoricalIndexConstituentsResponse = []HistoricalIndexConstituentChange

type HistoricalIndexConstituentChange struct {
	Date            types.Date `json:"date"`
	DateAdded       string     `json:"dateAdded"`
	Symbol          string     `json:"symbol"`
	AddedSecurity   string     `json:"addedSecurity"`
	RemovedTicker   *string    `json:"removedTicker"`
	RemovedSecurity *string    `json:"removedSecurity"`
	Reason          string     `json:"reason"`
}

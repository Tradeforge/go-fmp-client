package model

import (
	"go.tradeforge.dev/fmp/pkg/types"
)

type GetStockNewsParams struct {
	Symbols *string     `query:"tickers"`
	Since   *types.Date `query:"from"`
	Until   *types.Date `query:"to"`
	Page    *uint       `query:"page"`
	Limit   *uint       `query:"limit"`
}

type GetStockNewsResponse []NewsArticle

type NewsArticle struct {
	Symbol        string         `json:"symbol"`
	PublishedDate types.DateTime `json:"publishedDate"`
	Title         string         `json:"title"`
	Image         string         `json:"image"`
	Site          string         `json:"site"`
	Text          string         `json:"text"`
	URL           string         `json:"url"`
}

type GetNewsRSSFeedParams struct {
	Page uint `query:"page"`
}

type GetNewsRSSFeedResponse []NewsArticleWithSentiment

type NewsArticleWithSentiment struct {
	Symbol         string         `json:"symbol"`
	PublishedDate  types.DateTime `json:"publishedDate"`
	Title          string         `json:"title"`
	Image          string         `json:"image"`
	Site           string         `json:"site"`
	Text           string         `json:"text"`
	URL            string         `json:"url"`
	Sentiment      string         `json:"sentiment"`
	SentimentScore float64        `json:"sentimentScore"`
}

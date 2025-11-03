package model

import (
    "go.tradeforge.dev/fmp/pkg/types"
)

type GetFMPArticlesParams struct {
    Page  *uint `query:"page"`
    Limit *uint `query:"limit"`
}

type GetFMPArticlesResponse []FMPArticle

type FMPArticle struct {
    Ticker  string         `json:"ticker"`
    Date    types.DateTime `json:"date"`
    Title   string         `json:"title"`
    Image   string         `json:"image"`
    Site    string         `json:"site"`
    Content string         `json:"content"`
    Author  string         `json:"author"`
    Link    string         `json:"link"`
}

type GetNewsParams struct {
    Since *types.Date `query:"from"`
    Until *types.Date `query:"to"`
    Page  *uint       `query:"page"`
    Limit *uint       `query:"limit"`
}

type GetNewsResponse []NewsArticle

type NewsArticle struct {
    Symbol        string         `json:"symbol"`
    PublishedDate types.DateTime `json:"publishedDate"`
    Publisher     string         `json:"publisher"`
    Title         string         `json:"title"`
    Image         string         `json:"image"`
    Site          string         `json:"site"`
    Text          string         `json:"text"`
    URL           string         `json:"url"`
}

package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetStockNewsPath   = "/api/v3/stock_news"
	GetNewsRSSFeedPath = "/api/v4/stock-news-sentiments-rss-feed"
)

type NewsClient struct {
	*rest.Client
}

func (nc *NewsClient) GetStockNews(ctx context.Context, params model.GetStockNewsParams, opts ...model.RequestOption) (model.GetStockNewsResponse, error) {
	var res model.GetStockNewsResponse
	_, err := nc.Call(ctx, http.MethodGet, GetStockNewsPath, params, &res, opts...)
	return res, err
}

func (nc *NewsClient) GetNewsRSSFeed(ctx context.Context, params model.GetNewsRSSFeedParams, opts ...model.RequestOption) (model.GetNewsRSSFeedResponse, error) {
	var res model.GetNewsRSSFeedResponse
	_, err := nc.Call(ctx, http.MethodGet, GetNewsRSSFeedPath, params, &res, opts...)
	return res, err
}

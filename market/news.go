package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	ListStockNewsPath   = "/api/v3/stock_news"
	ListNewsRSSFeedPath = "/api/v4/stock-news-sentiments-rss-feed"
)

type NewsClient struct {
	*rest.Client
}

func (nc *NewsClient) ListStockNews(ctx context.Context, params model.ListStockNewsParams, opts ...model.RequestOption) (model.ListStockNewsResponse, error) {
	var res model.ListStockNewsResponse
	_, err := nc.Call(ctx, http.MethodGet, ListStockNewsPath, params, &res, opts...)
	return res, err
}

func (nc *NewsClient) ListNewsRSSFeed(ctx context.Context, params model.ListNewsRSSFeedParams, opts ...model.RequestOption) (model.ListNewsRSSFeedResponse, error) {
	var res model.ListNewsRSSFeedResponse
	_, err := nc.Call(ctx, http.MethodGet, ListNewsRSSFeedPath, params, &res, opts...)
	return res, err
}

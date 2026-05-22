package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetIndexListPath          = "/stable/index-list"
	BatchGetIndexQuotesPath   = "/stable/batch-index-quotes"
	GetHistoricalSP500Path    = "/stable/historical-sp500-constituent"
	GetHistoricalNasdaqPath   = "/stable/historical-nasdaq-constituent"
	GetHistoricalDowJonesPath = "/stable/historical-dowjones-constituent"
)

type IndexClient struct {
	*rest.Client
}

func (ic *IndexClient) GetIndexList(ctx context.Context, opts ...model.RequestOption) (model.GetIndexListResponse, error) {
	var res model.GetIndexListResponse
	_, err := ic.Call(ctx, http.MethodGet, GetIndexListPath, nil, &res, opts...)
	return res, err
}

// BatchGetIndexQuotes returns full quotes for all indexes.
func (ic *IndexClient) BatchGetIndexQuotes(ctx context.Context, opts ...model.RequestOption) (model.BatchGetIndexQuotesResponse, error) {
	var res model.BatchGetIndexQuotesResponse
	_, err := ic.Call(ctx, http.MethodGet, BatchGetIndexQuotesPath, nil, &res, append(opts, model.QueryParam("short", "false"))...)
	return res, err
}

// BatchGetIndexShortQuotes returns short quotes (symbol, price, change, volume) for all indexes.
func (ic *IndexClient) BatchGetIndexShortQuotes(ctx context.Context, opts ...model.RequestOption) (model.BatchGetIndexShortQuotesResponse, error) {
	var res model.BatchGetIndexShortQuotesResponse
	_, err := ic.Call(ctx, http.MethodGet, BatchGetIndexQuotesPath, nil, &res, opts...)
	return res, err
}

func (ic *IndexClient) GetHistoricalSP500Constituents(ctx context.Context, opts ...model.RequestOption) (model.GetHistoricalIndexConstituentsResponse, error) {
	var res model.GetHistoricalIndexConstituentsResponse
	_, err := ic.Call(ctx, http.MethodGet, GetHistoricalSP500Path, nil, &res, opts...)
	return res, err
}

func (ic *IndexClient) GetHistoricalNasdaqConstituents(ctx context.Context, opts ...model.RequestOption) (model.GetHistoricalIndexConstituentsResponse, error) {
	var res model.GetHistoricalIndexConstituentsResponse
	_, err := ic.Call(ctx, http.MethodGet, GetHistoricalNasdaqPath, nil, &res, opts...)
	return res, err
}

func (ic *IndexClient) GetHistoricalDowJonesConstituents(ctx context.Context, opts ...model.RequestOption) (model.GetHistoricalIndexConstituentsResponse, error) {
	var res model.GetHistoricalIndexConstituentsResponse
	_, err := ic.Call(ctx, http.MethodGet, GetHistoricalDowJonesPath, nil, &res, opts...)
	return res, err
}

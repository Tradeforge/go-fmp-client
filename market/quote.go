package market

import (
	"context"
	"fmt"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetQuotePath       = "/stable/quote"
	GetPriceChangePath = "/stable/stock-price-change"

	BatchGetQuotesPath           = "/stable/batch-quote"
	BatchGetQuotesByExchangePath = "/stable/batch-exchange-quote"

	GetHistoricalBarsPath      = "/stable/historical-chart/:timeframe"
	GetHistoricalPricesEODPath = "/stable/historical-price-eod/full"
	GetHistoricalMarketCapPath = "/stable/historical-market-capitalization"
)

type QuoteClient struct {
	*rest.Client
}

func (qc *QuoteClient) GetQuote(ctx context.Context, params *model.GetQuoteParams, opts ...model.RequestOption) (response *model.GetQuoteResponse, err error) {
	var res []model.GetQuoteResponse
	_, err = qc.Call(ctx, http.MethodGet, GetQuotePath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("expected response of length 1, got %d", len(res))
	}
	return &res[0], nil
}

func (qc *QuoteClient) BatchGetQuotes(ctx context.Context, params *model.BatchGetQuoteParams, opts ...model.RequestOption) (model.BatchGetQuoteResponse, error) {
	var res model.BatchGetQuoteResponse
	_, err := qc.Call(ctx, http.MethodGet, BatchGetQuotesPath, params, &res, opts...)
	return res, err
}

func (qc *QuoteClient) BatchGetQuotesByExchange(ctx context.Context, params *model.BatchGetQuotesByExchangeParams, opts ...model.RequestOption) (model.BatchGetQuotesByExchangeResponse, error) {
	params.Short = false // Ensure we always get full prices, not short ones.

	var res model.BatchGetQuotesByExchangeResponse
	_, err := qc.Call(ctx, http.MethodGet, BatchGetQuotesByExchangePath, params, &res, opts...)
	return res, err
}

func (qc *QuoteClient) GetPriceChange(ctx context.Context, params *model.GetPriceChangeParams, opts ...model.RequestOption) (response *model.GetPriceChangeResponse, err error) {
	var res []model.GetPriceChangeResponse
	_, err = qc.Call(ctx, http.MethodGet, GetPriceChangePath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("expected response of length 1, got %d", len(res))
	}
	return &res[0], nil
}

func (tc *TickerClient) GetHistoricalBars(ctx context.Context, params *model.GetHistoricalBarsParams, opts ...model.RequestOption) (model.GetHistoricalBarsResponse, error) {
	var res model.GetHistoricalBarsResponse
	_, err := tc.Call(ctx, http.MethodGet, GetHistoricalBarsPath, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetHistoricalPricesEOD(ctx context.Context, params *model.GetHistoricalPricesEODParams, opts ...model.RequestOption) (model.GetHistoricalPricesEODResponse, error) {
	var res model.GetHistoricalPricesEODResponse
	_, err := tc.Call(ctx, http.MethodGet, GetHistoricalPricesEODPath, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetHistoricalMarketCap(ctx context.Context, params *model.GetHistoricalMarketCapParams, opts ...model.RequestOption) (model.GetHistoricalMarketCapResponse, error) {
	var res model.GetHistoricalMarketCapResponse
	_, err := tc.Call(ctx, http.MethodGet, GetHistoricalMarketCapPath, params, &res, opts...)
	return res, err
}

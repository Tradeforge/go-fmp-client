package market

import (
	"context"
	"fmt"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	ListExchangeSymbolsPath = "/api/v3/symbol/:exchange"
	GetFullPricePath        = "/api/v3/quote/:symbol"
	GetPriceChangePath      = "/api/v3/stock-price-target/:symbol"
	BatchGetFullPricePath   = "/api/v3/quote/:symbols"

	GetRealtimeQuotePath      = "/api/v3/stock/full/real-time-price/:symbol"
	BatchGetRealtimeQuotePath = "/api/v3/stock/full/real-time-price/:symbols"
	ListAllRealtimeQuotesPath = "/api/v3/stock/full/real-time-price"
)

type QuoteClient struct {
	*rest.Client
}

func (qc *QuoteClient) GetFullPrice(ctx context.Context, params *model.GetFullPriceParams, opts ...model.RequestOption) (response *model.GetFullPriceResponse, err error) {
	var res []model.GetFullPriceResponse
	_, err = qc.Call(ctx, http.MethodGet, GetFullPricePath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("expected response of length 1, got %d", len(res))
	}
	return &res[0], nil
}

func (qc *QuoteClient) BatchGetFullPrice(ctx context.Context, params *model.BatchGetFullPriceParams, opts ...model.RequestOption) (model.BatchGetFullPriceResponse, error) {
	var res model.BatchGetFullPriceResponse
	_, err := qc.Call(ctx, http.MethodGet, BatchGetFullPricePath, params, &res, opts...)
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

func (qc *QuoteClient) BatchGetPriceChange(ctx context.Context, params *model.BatchGetPriceChangeParams, opts ...model.RequestOption) ([]model.GetPriceChangeResponse, error) {
	var res []model.GetPriceChangeResponse
	_, err := qc.Call(ctx, http.MethodGet, BatchGetRealtimeQuotePath, params, &res, opts...)
	return res, err
}

func (qc *QuoteClient) GetRealtimeQuote(ctx context.Context, params *model.GetRealtimeQuoteParams, opts ...model.RequestOption) (response *model.GetRealtimeQuoteResponse, err error) {
	var res []model.GetRealtimeQuoteResponse
	_, err = qc.Call(ctx, http.MethodGet, GetRealtimeQuotePath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("expected response of length 1, got %d", len(res))
	}
	return &res[0], nil
}

func (qc *QuoteClient) BatchGetRealtimeQuote(ctx context.Context, params *model.BatchGetRealtimeQuoteParams, opts ...model.RequestOption) (model.BatchGetRealtimeQuoteResponse, error) {
	var res model.BatchGetRealtimeQuoteResponse
	_, err := qc.Call(ctx, http.MethodGet, BatchGetRealtimeQuotePath, params, &res, opts...)
	return res, err
}

func (qc *QuoteClient) ListAllRealtimeQuotes(ctx context.Context, opts ...model.RequestOption) (model.ListAllRealtimeQuotesResponse, error) {
	var res model.ListAllRealtimeQuotesResponse
	_, err := qc.Call(ctx, http.MethodGet, ListAllRealtimeQuotesPath, nil, &res, opts...)
	return res, err
}

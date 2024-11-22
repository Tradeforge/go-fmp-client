package market

import (
	"context"
	"fmt"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetRealTimeQuotePath = "/api/v3/stock/full/real-time-price/:symbol"
	GetFullPricePath     = "/api/v3/quote/:symbol"
	GetPriceChangePath   = "/api/v3/stock-price-target/:symbol"

	BatchGetRealTimeQuotePath = "/api/v3/stock/full/real-time-price/:symbols"
	BatchGetFullPricePath     = "/api/v3/quote/:symbols"
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
	_, err := qc.Call(ctx, http.MethodGet, BatchGetRealTimeQuotePath, params, &res, opts...)
	return res, err
}

func (qc *QuoteClient) GetRealTimeQuote(ctx context.Context, params *model.GetRealTimeQuoteParams, opts ...model.RequestOption) (response *model.GetRealTimeQuoteResponse, err error) {
	var res []model.GetRealTimeQuoteResponse
	_, err = qc.Call(ctx, http.MethodGet, GetRealTimeQuotePath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("expected response of length 1, got %d", len(res))
	}
	return &res[0], nil
}

func (qc *QuoteClient) BatchGetRealTimeQuote(ctx context.Context, params *model.BatchGetRealTimeQuoteParams, opts ...model.RequestOption) (model.BatchGetRealTimeQuoteResponse, error) {
	var res model.BatchGetRealTimeQuoteResponse
	_, err := qc.Call(ctx, http.MethodGet, BatchGetRealTimeQuotePath, params, &res, opts...)
	return res, err
}

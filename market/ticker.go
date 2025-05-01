package market

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	ListMostActiveTickersPath = "/api/v3/stock_market/actives"
	ListGainersPath           = "/api/v3/stock_market/gainers"
	ListLosersPath            = "/api/v3/stock_market/losers"

	GetCompanyProfilePath     = "/api/v3/profile/:symbol"
	ListCompanyProfilesPath   = "/api/v3/profile/:symbols"
	BulkGetCompanyProfilePath = "/stable/profile-bulk"

	ListTickerKeyMetricsPath = "/api/v3/key-metrics/:symbol"
	ListTickerRatiosPath     = "/api/v3/ratios/:symbol"
)

type TickerClient struct {
	*rest.Client
}

func (tc *TickerClient) GetCompanyProfile(ctx context.Context, params *model.GetCompanyProfileParams, opts ...model.RequestOption) (_ *model.GetCompanyProfileResponse, err error) {
	var res []model.GetCompanyProfileResponse
	_, err = tc.Call(ctx, http.MethodGet, GetCompanyProfilePath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		// Return empty response if no data is returned.
		// Since we're using FMP for the stock profile but the stock assets are managed by Alpaca,
		// it might happen that some stocks are not available in FMP. For that reason, we don't want to
		// return an error if no data is found.
		return &model.GetCompanyProfileResponse{
			Symbol: params.Symbol,
		}, nil
	}
	return &res[0], nil
}

func (tc *TickerClient) ListCompanyProfiles(ctx context.Context, params *model.ListCompanyProfilesParams, opts ...model.RequestOption) ([]model.GetCompanyProfileResponse, error) {
	var res []model.GetCompanyProfileResponse
	_, err := tc.Call(ctx, http.MethodGet, ListCompanyProfilesPath, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) BulkGetCompanyProfile(ctx context.Context, params *model.BulkGetCompanyProfilesParams, opts ...model.RequestOption) ([]model.BulkCompanyProfileResponse, error) {
	r, err := tc.Call(
		ctx,
		http.MethodGet,
		BulkGetCompanyProfilePath,
		params,
		// No response means the original response will be returned as is.
		nil,
		// We need to ignore the bad request status code because the API returns a 400 status code when there is no more data to fetch.
		append(opts, model.WithContentType("text/csv"), model.WithIgnoredErrorStatusCodes(http.StatusBadRequest))...)
	if err != nil {
		return nil, err
	}
	if r.StatusCode() == http.StatusBadRequest {
		// Return empty response if no data is returned. FMP sends a 400 status code when there is no more data to fetch.
		return []model.BulkCompanyProfileResponse{}, nil
	}
	csvReader := csv.NewReader(bytes.NewReader(r.Body()))
	h, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("reading header: %w", err)
	}
	d, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("reading records: %w", err)
	}
	res := make([]model.BulkCompanyProfileResponse, 0, len(d))
	for _, record := range d {
		if len(record) != len(h) {
			return nil, fmt.Errorf("invalid record length: expected %d, got %d", len(h), len(record))
		}
		p, err := model.ParseCompanyProfileCSVRecord(h, record)
		if err != nil {
			return nil, fmt.Errorf("parsing record: %w", err)
		}
		res = append(res, *p)
	}

	return res, err
}

func (tc *TickerClient) ListTickerKeyMetrics(ctx context.Context, params *model.ListStockKeyMetricsParams, opts ...model.RequestOption) (model.ListTickerKeyMetricsResponse, error) {
	var res model.ListTickerKeyMetricsResponse
	_, err := tc.Call(ctx, http.MethodGet, ListTickerKeyMetricsPath, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) ListTickerRatios(ctx context.Context, params *model.ListStockRatiosParams, opts ...model.RequestOption) (model.ListTickerRatiosResponse, error) {
	var res model.ListTickerRatiosResponse
	_, err := tc.Call(ctx, http.MethodGet, ListTickerRatiosPath, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) ListGainers(ctx context.Context, opts ...model.RequestOption) (model.ListGainersResponse, error) {
	var res model.ListGainersResponse
	_, err := tc.Call(ctx, http.MethodGet, ListGainersPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) ListLosers(ctx context.Context, opts ...model.RequestOption) (model.ListLosersResponse, error) {
	var res model.ListLosersResponse
	_, err := tc.Call(ctx, http.MethodGet, ListLosersPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) ListMostActiveTickers(ctx context.Context, opts ...model.RequestOption) (model.ListMostActiveTickersResponse, error) {
	var res model.ListMostActiveTickersResponse
	_, err := tc.Call(ctx, http.MethodGet, ListMostActiveTickersPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) ListExchangeSymbols(ctx context.Context, params *model.ListExchangeSymbolsParams, opts ...model.RequestOption) (model.ListExchangeSymbolsResponse, error) {
	var res model.ListExchangeSymbolsResponse
	_, err := tc.Call(ctx, http.MethodGet, ListExchangeSymbolsPath, params, &res, opts...)
	return res, err
}

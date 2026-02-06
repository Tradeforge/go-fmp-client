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
	GetMostActiveTickersPath = "/stable/most-actives"
	GetGainersPath           = "/stable/biggest-gainers"
	GetLosersPath            = "/stable/biggest-losers"

	GetCompanyProfilePath     = "/stable/profile"
	BulkGetCompanyProfilePath = "/stable/profile-bulk"

	GetFinancialKeyMetricsTTMPath = "/stable/key-metrics-ttm"
	GetFinancialRatiosTTMPath     = "/stable/ratios-ttm"
	GetIncomeStatements           = "/stable/income-statement"
	GetBalanceSheets              = "/stable/balance-sheet-statement"
	GetCashFlowStatements         = "/stable/cash-flow-statement"

	GetSP500IndexConstituentsPath    = "/stable/sp500-constituent"
	GetNasdaqIndexConstituentsPath   = "/stable/nasdaq-constituent"
	GetDowJonesIndexConstituentsPath = "/stable/dowjones-constituent"

	GetAvailableExchangesPath = "/stable/available-exchanges"
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

func (tc *TickerClient) GetFinancialKeyMetricsTTM(ctx context.Context, params *model.GetFinancialKeyMetricsTTMParams, opts ...model.RequestOption) (model.GetFinancialKeyMetricsTTMResponse, error) {
	var res []model.GetFinancialKeyMetricsTTMResponse
	_, err := tc.Call(ctx, http.MethodGet, GetFinancialKeyMetricsTTMPath, params, &res, opts...)
	if len(res) == 0 {
		return nil, err
	}
	return res[0], err
}

func (tc *TickerClient) GetFinancialRatiosTTM(ctx context.Context, params *model.GetFinancialRatiosTTMParams, opts ...model.RequestOption) (model.GetFinancialRatiosTTMResponse, error) {
	var res []model.GetFinancialRatiosTTMResponse
	_, err := tc.Call(ctx, http.MethodGet, GetFinancialRatiosTTMPath, params, &res, opts...)
	if len(res) == 0 {
		return nil, err
	}
	return res[0], err
}

func (tc *TickerClient) GetIncomeStatements(ctx context.Context, params *model.GetIncomeStatementsParams, opts ...model.RequestOption) (model.GetIncomeStatementsResponse, error) {
	var res model.GetIncomeStatementsResponse
	_, err := tc.Call(ctx, http.MethodGet, GetIncomeStatements, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetBalanceSheets(ctx context.Context, params *model.GetBalanceSheetsParams, opts ...model.RequestOption) (model.GetBalanceSheetsResponse, error) {
	var res model.GetBalanceSheetsResponse
	_, err := tc.Call(ctx, http.MethodGet, GetBalanceSheets, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetCashFlowStatements(ctx context.Context, params *model.GetCashFlowStatementsParams, opts ...model.RequestOption) (model.GetCashFlowStatementsResponse, error) {
	var res model.GetCashFlowStatementsResponse
	_, err := tc.Call(ctx, http.MethodGet, GetCashFlowStatements, params, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetGainers(ctx context.Context, opts ...model.RequestOption) (model.GetGainersResponse, error) {
	var res model.GetGainersResponse
	_, err := tc.Call(ctx, http.MethodGet, GetGainersPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetLosers(ctx context.Context, opts ...model.RequestOption) (model.GetLosersResponse, error) {
	var res model.GetLosersResponse
	_, err := tc.Call(ctx, http.MethodGet, GetLosersPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetMostActiveTickers(ctx context.Context, opts ...model.RequestOption) (model.GetMostActiveTickersResponse, error) {
	var res model.GetMostActiveTickersResponse
	_, err := tc.Call(ctx, http.MethodGet, GetMostActiveTickersPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetSP500IndexConstituents(ctx context.Context, opts ...model.RequestOption) (model.GetIndexConstituentsResponse, error) {
	var res model.GetIndexConstituentsResponse
	_, err := tc.Call(ctx, http.MethodGet, GetSP500IndexConstituentsPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetNasdaqIndexConstituents(ctx context.Context, opts ...model.RequestOption) (model.GetIndexConstituentsResponse, error) {
	var res model.GetIndexConstituentsResponse
	_, err := tc.Call(ctx, http.MethodGet, GetNasdaqIndexConstituentsPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetDowJonesIndexConstituents(ctx context.Context, opts ...model.RequestOption) (model.GetIndexConstituentsResponse, error) {
	var res model.GetIndexConstituentsResponse
	_, err := tc.Call(ctx, http.MethodGet, GetDowJonesIndexConstituentsPath, nil, &res, opts...)
	return res, err
}

func (tc *TickerClient) GetAvailableExchanges(ctx context.Context, opts ...model.RequestOption) (model.GetAvailableExchangesResponse, error) {
	var res model.GetAvailableExchangesResponse
	_, err := tc.Call(ctx, http.MethodGet, GetAvailableExchangesPath, nil, &res, opts...)
	return res, err
}

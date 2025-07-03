package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetEarningsCalendarPath   = "/api/v3/earning_calendar"
	GetHistoricalEarningsPath = "/api/v3/historical/earning_calendar/:symbol"
	GetSECFilingsRSSFeedPath  = "/api/v4/rss_feed"
	GetInsiderTradesPath      = "/stable/insider-trading/latest"

	defaultSECFeedLimit = 100
)

type EventClient struct {
	*rest.Client
}

func (ec *EventClient) GetInsiderTrades(ctx context.Context, params model.GetInsiderTradesParams, opts ...model.RequestOption) (model.GetInsiderTradesResponse, error) {
	var res model.GetInsiderTradesResponse
	_, err := ec.Call(ctx, http.MethodGet, GetInsiderTradesPath, params, &res, opts...)
	return res, err
}

func (ec *EventClient) GetSECFilingsRSSFeed(ctx context.Context, params model.GetSECFilingsRSSFeedParams, opts ...model.RequestOption) (model.GetSECFilingsRSSFeedResponse, error) {
	r, err := ec.doGetSECFilingsRSSFeed(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	filteredResults := make(model.GetSECFilingsRSSFeedResponse, 0, len(r))
L:
	for {
		for _, filing := range r {
			if filing.Type.Form == params.Type {
				filteredResults = append(filteredResults, filing)
			}
			if len(filteredResults) == len(r) {
				break L
			}
		}
		r, err = ec.doGetSECFilingsRSSFeed(ctx, params, opts...)
		if err != nil {
			return nil, err
		}
	}
	return filteredResults, nil
}

func (ec *EventClient) GetEarningsCalendar(ctx context.Context, params *model.GetEarningsCalendarParams, opts ...model.RequestOption) ([]model.GetEarningsCalendarResponse, error) {
	var res []model.GetEarningsCalendarResponse
	_, err := ec.Call(ctx, http.MethodGet, GetEarningsCalendarPath, params, &res, opts...)
	return res, err
}

func (ec *EventClient) GetHistoricalEarningsCalendar(ctx context.Context, params *model.GetHistoricalEarningsCalendarParams, opts ...model.RequestOption) ([]model.GetEarningsCalendarResponse, error) {
	var res []model.GetEarningsCalendarResponse
	_, err := ec.Call(ctx, http.MethodGet, GetHistoricalEarningsPath, params, &res, opts...)
	return res, err
}

func (ec *EventClient) doGetSECFilingsRSSFeed(ctx context.Context, params model.GetSECFilingsRSSFeedParams, opts ...model.RequestOption) (model.GetSECFilingsRSSFeedResponse, error) {
	var limit uint = defaultSECFeedLimit
	if params.Limit == nil {
		params.Limit = &limit
	}
	var res model.GetSECFilingsRSSFeedResponse
	_, err := ec.Call(ctx, http.MethodGet, GetSECFilingsRSSFeedPath, params, &res, opts...)
	return res, err
}

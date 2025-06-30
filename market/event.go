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
    ListSECFilingsRSSFeedPath = "/api/v4/rss_feed"
    ListInsiderTradesPath     = "/stable/insider-trading/latest"

    defaultSECFeedLimit = 100
)

type EventClient struct {
    *rest.Client
}

func (ec *EventClient) ListInsiderTrades(ctx context.Context, params model.ListInsiderTradesParams, opts ...model.RequestOption) (model.ListInsiderTradesResponse, error) {
    var res model.ListInsiderTradesResponse
    _, err := ec.Call(ctx, http.MethodGet, ListInsiderTradesPath, params, &res, opts...)
    return res, err
}

func (ec *EventClient) ListSECFilingsRSSFeed(ctx context.Context, params model.ListSECFilingsRSSFeedParams, opts ...model.RequestOption) (model.ListSECFilingsRSSFeedResponse, error) {
    r, err := ec.doListSECFilingsRSSFeed(ctx, params, opts...)
    if err != nil {
        return nil, err
    }
    filteredResults := make(model.ListSECFilingsRSSFeedResponse, 0, len(r))
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
        r, err = ec.doListSECFilingsRSSFeed(ctx, params, opts...)
        if err != nil {
            return nil, err
        }
    }
    return filteredResults, nil
}

func (ec *EventClient) doListSECFilingsRSSFeed(ctx context.Context, params model.ListSECFilingsRSSFeedParams, opts ...model.RequestOption) (model.ListSECFilingsRSSFeedResponse, error) {
    var limit uint = defaultSECFeedLimit
    if params.Limit == nil {
        params.Limit = &limit
    }
    var res model.ListSECFilingsRSSFeedResponse
    _, err := ec.Call(ctx, http.MethodGet, ListSECFilingsRSSFeedPath, params, &res, opts...)
    return res, err
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

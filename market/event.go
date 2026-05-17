package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetEarningsCalendarPath  = "/stable/earnings-calendar"
	GetDividendsCalendarPath = "/stable/dividends-calendar"
	GetInsiderTradesPath     = "/stable/insider-trading/latest"
)

type EventClient struct {
	*rest.Client
}

func (ec *EventClient) GetInsiderTrades(ctx context.Context, params model.GetInsiderTradesParams, opts ...model.RequestOption) (model.GetInsiderTradesResponse, error) {
	var res model.GetInsiderTradesResponse
	_, err := ec.Call(ctx, http.MethodGet, GetInsiderTradesPath, params, &res, opts...)
	return res, err
}

func (ec *EventClient) GetEarningsCalendar(ctx context.Context, params *model.GetEarningsCalendarParams, opts ...model.RequestOption) ([]model.GetEarningsCalendarResponse, error) {
	var res []model.GetEarningsCalendarResponse
	_, err := ec.Call(ctx, http.MethodGet, GetEarningsCalendarPath, params, &res, opts...)
	return res, err
}

func (ec *EventClient) GetDividendsCalendar(ctx context.Context, params *model.GetDividendsCalendarParams, opts ...model.RequestOption) ([]model.GetDividendsCalendarResponse, error) {
	var res []model.GetDividendsCalendarResponse
	_, err := ec.Call(ctx, http.MethodGet, GetDividendsCalendarPath, params, &res, opts...)
	return res, err
}

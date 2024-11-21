package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetEarningsCalendarPath   = "/api/v3/earning_calendar"
	GetHistoricalEarningsPath = "/api/v3/historical/earning_calendar/:symbol"
)

type EventClient struct {
	*client.Client
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

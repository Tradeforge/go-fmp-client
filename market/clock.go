package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetAllExchangesTradingHoursPath = "/api/v3/is-the-market-open-all"
	GetExchangeHolidaysPath         = "/api/v3/is-the-market-open"
)

type ClockClient struct {
	*rest.Client
}

func (cc *ClockClient) GetAllExchangesTradingHours(ctx context.Context, opts ...model.RequestOption) (model.GetAllExchangesTradingHoursResponse, error) {
	var res model.GetAllExchangesTradingHoursResponse
	_, err := cc.Call(ctx, http.MethodGet, GetAllExchangesTradingHoursPath, nil, &res, opts...)
	return res, err
}

func (cc *ClockClient) GetExchangeHolidays(ctx context.Context, params model.GetExchangeHolidaysParams, opts ...model.RequestOption) (model.GetExchangeHolidaysResponse, error) {
	var aux model.OriginalGetExchangeHolidaysResponse
	_, err := cc.Call(ctx, http.MethodGet, GetExchangeHolidaysPath, params, &aux, opts...)

	var res model.GetExchangeHolidaysResponse
	for _, holiday := range aux.HolidaysByYear {
		if holiday.Year() != params.Year {
			continue
		}
		for _, h := range holiday.Get() {
			res = append(res, model.Holiday{
				Date: h.Date,
				Name: h.Name,
			})
		}
	}
	return res, err
}

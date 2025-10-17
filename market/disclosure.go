package market

import (
	"context"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const GetHouseFinancialDisclosuresPath = "/stable/house-latest"

type DisclosureClient struct {
	*rest.Client
}

func (dc *DisclosureClient) GetHouseFinancialDisclosures(ctx context.Context, params model.GetHouseFinancialDisclosuresParams, opts ...model.RequestOption) (model.GetHouseFinancialDisclosuresResponse, error) {
	var res model.GetHouseFinancialDisclosuresResponse
	_, err := dc.Call(ctx, http.MethodGet, GetHouseFinancialDisclosuresPath, params, &res, opts...)
	return res, err
}

func (dc *DisclosureClient) GetSenateFinancialDisclosures(ctx context.Context, params model.GetSenateFinancialDisclosuresParams, opts ...model.RequestOption) (model.GetSenateFinancialDisclosuresResponse, error) {
	var res model.GetSenateFinancialDisclosuresResponse
	_, err := dc.Call(ctx, http.MethodGet, GetHouseFinancialDisclosuresPath, params, &res, opts...)
	return res, err
}

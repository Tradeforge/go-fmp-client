package market

import (
	"context"
	"errors"
	"net/http"

	"go.tradeforge.dev/fmp/client/rest"
	"go.tradeforge.dev/fmp/model"
)

const (
	GetAdvancedDCFPath = "/stable/custom-discounted-cash-flow"
)

type AnalysisClient struct {
	*rest.Client
}

func (ac *AnalysisClient) GetAdvancedDCF(ctx context.Context, params *model.GetAdvancedDCFParams, opts ...model.RequestOption) ([]model.GetAdvancedDCFResponse, error) {
	var res []model.GetAdvancedDCFResponse
	_, err := ac.Call(ctx, http.MethodGet, GetAdvancedDCFPath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("expected non-empty response, got 0 results")
	}
	return res, nil
}

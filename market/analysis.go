package market

import (
	"context"
	"fmt"
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

func (ac *AnalysisClient) GetAdvancedDCF(ctx context.Context, params *model.GetAdvancedDCFParams, opts ...model.RequestOption) (*model.GetAdvancedDCFResponse, error) {
	var res []model.GetAdvancedDCFResponse
	_, err := ac.Call(ctx, http.MethodGet, GetAdvancedDCFPath, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	if len(res) != 1 {
		return nil, fmt.Errorf("expected response of length 1, got %d", len(res))
	}
	return &res[0], nil
}
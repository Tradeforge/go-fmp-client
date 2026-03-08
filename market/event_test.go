package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
	"go.tradeforge.dev/fmp/pkg/types"
)

func TestGetInsiderTrades(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetInsiderTrades(ctx, model.GetInsiderTradesParams{})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, trade := range res[:5] {
		assert.NotEmpty(t, trade.Symbol)
		assert.NotEmpty(t, string(trade.TransactionDate))
		assert.NotEmpty(t, trade.OwnerName)
		assert.NotEmpty(t, string(trade.FilingDate))
	}
}

func TestGetEarningsCalendar(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	since := types.Date("2025-01-06")
	until := types.Date("2025-01-10")
	res, err := client.GetEarningsCalendar(ctx, &model.GetEarningsCalendarParams{
		Since: &since,
		Until: &until,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, e := range res[:5] {
		assert.NotEmpty(t, e.Symbol)
		assert.NotEmpty(t, string(e.Date))
	}
}

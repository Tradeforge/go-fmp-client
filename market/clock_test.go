package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
)

func TestGetAllExchangesTradingHours(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetAllExchangesTradingHours(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, eth := range res {
		assert.NotEmpty(t, eth.Exchange, "Exchange name should not be empty")
		assert.NotEmpty(t, eth.OpeningHour, "OpeningHour should not be empty")
		assert.NotEmpty(t, eth.ClosingHour, "ClosingHour should not be empty")
		assert.NotEmpty(t, eth.TimeZone, "TimeZone should not be empty")
	}
}

func TestGetExchangeHolidays(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetExchangeHolidays(ctx, model.GetExchangeHolidaysParams{
		Exchange: "NYSE",
		Year:     2025,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, h := range res {
		assert.NotEmpty(t, h.Date, "Holiday date should not be empty")
		assert.NotEmpty(t, h.Name, "Holiday name should not be empty")
	}
}

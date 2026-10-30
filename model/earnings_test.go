package model

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/pkg/types"
)

// The fixture uses the key set /stable/earnings-calendar returned on
// 2026-10-30 (symbol, date, epsActual, epsEstimated, revenueActual,
// revenueEstimated, lastUpdated); the values are made up.
//
//go:embed fixtures/earnings_calendar.json
var rawEarningsCalendar string

func TestGetEarningsCalendarResponse_UnmarshalJSON(t *testing.T) {
	var earnings []GetEarningsCalendarResponse
	require.NoError(t, json.Unmarshal([]byte(rawEarningsCalendar), &earnings))
	require.Len(t, earnings, 2)

	reported := earnings[0]
	assert.Equal(t, "AAPL", reported.Symbol)
	assert.Equal(t, types.Date("2026-10-29"), reported.Date)
	assert.Equal(t, types.Date("2026-10-30"), reported.LastUpdatedAt)
	require.NotNil(t, reported.EPS, "epsActual must decode into EPS")
	assert.True(t, decimal.RequireFromString("1.85").Equal(*reported.EPS))
	require.NotNil(t, reported.Revenue, "revenueActual must decode into Revenue")
	assert.True(t, decimal.RequireFromString("102466000000").Equal(*reported.Revenue))
	require.NotNil(t, reported.EPSEstimated)
	assert.True(t, decimal.RequireFromString("1.77").Equal(*reported.EPSEstimated))
	require.NotNil(t, reported.RevenueEstimated)
	assert.True(t, decimal.RequireFromString("101850000000").Equal(*reported.RevenueEstimated))

	upcoming := earnings[1]
	assert.Nil(t, upcoming.EPS)
	assert.Nil(t, upcoming.Revenue)
	require.NotNil(t, upcoming.EPSEstimated)
	assert.True(t, decimal.RequireFromString("3.67").Equal(*upcoming.EPSEstimated))
}

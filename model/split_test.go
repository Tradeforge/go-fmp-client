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

//go:embed fixtures/splits_calendar.json
var rawSplitsCalendar string

func TestGetSplitsCalendarResponse_UnmarshalJSON(t *testing.T) {
	var splits []GetSplitsCalendarResponse
	err := json.Unmarshal([]byte(rawSplitsCalendar), &splits)
	require.NoError(t, err)
	require.Len(t, splits, 4)

	for _, s := range splits {
		assert.NotEmpty(t, s.Symbol)
		assert.NotEmpty(t, string(s.Date))
	}

	// Forward split: 2:1.
	forward := splits[0]
	assert.Equal(t, "SFBS", forward.Symbol)
	assert.Equal(t, types.Date("2026-08-21"), forward.Date)
	assert.True(t, decimal.NewFromInt(2).Equal(forward.Numerator))
	assert.True(t, decimal.NewFromInt(1).Equal(forward.Denominator))
	assert.Equal(t, "stock-split", forward.SplitType)

	// Reverse split: 1:5.
	reverse := splits[1]
	assert.True(t, decimal.NewFromInt(1).Equal(reverse.Numerator))
	assert.True(t, decimal.NewFromInt(5).Equal(reverse.Denominator))

	// Stock dividend rides the same endpoint with a distinct split type.
	stockDividend := splits[2]
	assert.Equal(t, "stock-dividend", stockDividend.SplitType)
	assert.True(t, decimal.NewFromInt(51).Equal(stockDividend.Numerator))
	assert.True(t, decimal.NewFromInt(50).Equal(stockDividend.Denominator))

	// Missing split type decodes to empty rather than erroring.
	untyped := splits[3]
	assert.Empty(t, untyped.SplitType)
}

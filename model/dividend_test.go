package model

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/pkg/types"
)

//go:embed fixtures/dividends_calendar.json
var rawDividendsCalendar string

func TestGetDividendsCalendarResponse_UnmarshalJSON(t *testing.T) {
	var dividends []GetDividendsCalendarResponse
	err := json.Unmarshal([]byte(rawDividendsCalendar), &dividends)
	require.NoError(t, err)
	require.Len(t, dividends, 3)

	for _, d := range dividends {
		assert.NotEmpty(t, d.Symbol)
		assert.NotEmpty(t, string(d.Date))
	}

	// Fully-populated record: the secondary dates decode to present values.
	full := dividends[0]
	assert.False(t, full.RecordDate.IsEmpty())
	assert.False(t, full.PaymentDate.IsEmpty())
	assert.False(t, full.DeclarationDate.IsEmpty())
	require.NotNil(t, full.RecordDate.Value())
	assert.Equal(t, types.Date("2025-02-10"), *full.RecordDate.Value())

	// Record with empty "" secondary dates: EmptyOr decodes them as empty
	// instead of erroring on time.Parse of an empty string.
	empty := dividends[1]
	assert.True(t, empty.RecordDate.IsEmpty())
	assert.True(t, empty.PaymentDate.IsEmpty())
	assert.True(t, empty.DeclarationDate.IsEmpty())
}

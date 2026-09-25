package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/pkg/types"
)

// TestTickerQuote_UnmarshalJSON_Timestamp pins BE-476: FMP started serialising
// the quote timestamp as a fractional number, which made every
// syncExchangePrices invocation fail to unmarshal. Both the fractional and the
// integral form have to decode, and the decoded value has to be the whole
// second the quote fell in.
func TestTickerQuote_UnmarshalJSON_Timestamp(t *testing.T) {
	tests := []struct {
		name string
		// rawTimestamp is spliced into the quote payload verbatim, so each case
		// exercises the real struct rather than the type in isolation.
		rawTimestamp string
		want         types.UnixTimestamp
		wantErrMsg   string
	}{
		{
			name:         "success:integral-epoch",
			rawTimestamp: `1768692032`,
			want:         1768692032,
		},
		// The three fractional payloads below are verbatim production values
		// taken from the Sentry issue (TRADEFORGE-API-J).
		{
			name:         "success:fractional-epoch-three-decimals",
			rawTimestamp: `1768692032.131`,
			want:         1768692032,
		},
		{
			name:         "success:fractional-epoch-three-decimals-rounding-up",
			rawTimestamp: `1768852017.738`,
			want:         1768852017,
		},
		{
			name:         "success:fractional-epoch-two-decimals",
			rawTimestamp: `1768793276.38`,
			want:         1768793276,
		},
		{
			name:         "success:zero-epoch",
			rawTimestamp: `0`,
			want:         0,
		},
		{
			name:         "success:fractional-below-one-second",
			rawTimestamp: `0.999`,
			want:         0,
		},
		{
			name:         "success:negative-integral-epoch",
			rawTimestamp: `-86400`,
			want:         -86400,
		},
		{
			// Flooring, not truncation toward zero: -86399.5 sits inside the
			// second that starts at -86400, so that is the second it decodes
			// to. Truncating would push it forward into the next one.
			name:         "success:negative-fractional-epoch",
			rawTimestamp: `-86399.5`,
			want:         -86400,
		},
		{
			// The plain int64 this type replaced also decoded null to zero;
			// keeping that avoids turning a sparse field into a batch failure.
			name:         "success:null-decodes-to-zero",
			rawTimestamp: `null`,
			want:         0,
		},
		{
			name:         "failure:numeric-string",
			rawTimestamp: `"1768692032"`,
			wantErrMsg:   "invalid unix timestamp: expected a number, got a string",
		},
		{
			name:         "failure:non-numeric-string",
			rawTimestamp: `"not-a-timestamp"`,
			wantErrMsg:   "invalid unix timestamp: expected a number, got a string",
		},
		{
			name:         "failure:boolean",
			rawTimestamp: `true`,
			wantErrMsg:   "invalid unix timestamp",
		},
		{
			name:         "failure:object",
			rawTimestamp: `{"seconds":1768692032}`,
			wantErrMsg:   "invalid unix timestamp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			raw := `{"symbol":"AAPL","name":"Apple Inc.","exchange":"NASDAQ","timestamp":` + tt.rawTimestamp + `}`

			var quote TickerQuote
			err := json.Unmarshal([]byte(raw), &quote)
			if tt.wantErrMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErrMsg)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, "AAPL", quote.Symbol)
			assert.Equal(t, tt.want, quote.Timestamp)
			assert.Equal(t, int64(tt.want), quote.Timestamp.Int64())
			assert.Equal(t, time.Unix(int64(tt.want), 0), quote.Timestamp.Time())
		})
	}
}

// TestTickerQuote_MarshalJSON_TimestampRoundTrip proves a decoded fractional
// timestamp is re-emitted as the whole second, so anything that persists or
// forwards a TickerQuote keeps the wire format the plain int64 produced.
func TestTickerQuote_MarshalJSON_TimestampRoundTrip(t *testing.T) {
	var quote TickerQuote
	require.NoError(t, json.Unmarshal([]byte(`{"symbol":"AAPL","timestamp":1768692032.131}`), &quote))

	encoded, err := json.Marshal(quote)
	require.NoError(t, err)
	assert.Contains(t, string(encoded), `"timestamp":1768692032`)

	var decoded TickerQuote
	require.NoError(t, json.Unmarshal(encoded, &decoded))
	assert.Equal(t, quote.Timestamp, decoded.Timestamp)
}

// TestUnixTimestamp_Time proves the decoded second is the one the fractional
// payload named, rather than merely a value that parsed.
func TestUnixTimestamp_Time(t *testing.T) {
	var quote TickerQuote
	require.NoError(t, json.Unmarshal([]byte(`{"timestamp":1768692032.131}`), &quote))

	assert.Equal(t, "2026-01-17T23:20:32Z", quote.Timestamp.Time().UTC().Format(time.RFC3339))
	assert.Equal(t, "1768692032", quote.Timestamp.String())
	assert.Equal(t, quote.Timestamp, types.UnixTimestampFromTime(quote.Timestamp.Time()))
}

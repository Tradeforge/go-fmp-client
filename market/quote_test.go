package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
	"go.tradeforge.dev/fmp/pkg/types"
)

func TestGetQuote(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetQuote(ctx, &model.GetQuoteParams{Symbol: "AAPL"})
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, "AAPL", res.Symbol)
	assert.NotEmpty(t, res.Name)
	assert.True(t, res.Price.IsPositive(), "price should be positive")
	assert.NotEmpty(t, res.Exchange)
	assert.NotZero(t, res.Timestamp)
}

func TestGetPriceChange(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetPriceChange(ctx, &model.GetPriceChangeParams{Symbol: "AAPL"})
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, "AAPL", res.Symbol)
}

func TestBatchGetQuotes(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.BatchGetQuotes(ctx, &model.BatchGetQuoteParams{Symbols: "AAPL,MSFT"})
	require.NoError(t, err)
	require.Len(t, res, 2)

	symbols := map[string]bool{}
	for _, q := range res {
		symbols[q.Symbol] = true
		assert.NotEmpty(t, q.Name)
		assert.True(t, q.Price.IsPositive())
	}
	assert.True(t, symbols["AAPL"])
	assert.True(t, symbols["MSFT"])
}

func TestBatchGetQuotesByExchange(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.BatchGetQuotesByExchange(ctx, &model.BatchGetQuotesByExchangeParams{Exchange: "NYSE"})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, q := range res[:5] {
		assert.NotEmpty(t, q.Symbol)
		assert.NotEmpty(t, q.Exchange)
	}
}

func TestGetHistoricalBars(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetHistoricalBars(ctx, &model.GetHistoricalBarsParams{
		Timeframe: model.Timeframe1Hour,
		Symbol:    "AAPL",
		Since:     types.Date("2025-01-02"),
		Until:     types.Date("2025-01-03"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, bar := range res {
		assert.True(t, bar.Open.IsPositive())
		assert.True(t, bar.High.IsPositive())
		assert.True(t, bar.Low.IsPositive())
		assert.True(t, bar.Close.IsPositive())
		assert.NotEmpty(t, string(bar.DateTime))
	}
}

func TestGetHistoricalPricesEOD(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetHistoricalPricesEOD(ctx, &model.GetHistoricalPricesEODParams{
		Symbol: "AAPL",
		Since:  types.Date("2025-01-02"),
		Until:  types.Date("2025-01-10"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, p := range res {
		assert.Equal(t, "AAPL", p.Symbol)
		assert.NotEmpty(t, string(p.Date))
		assert.True(t, p.Open.IsPositive())
		assert.True(t, p.High.IsPositive())
		assert.True(t, p.Low.IsPositive())
		assert.True(t, p.Close.IsPositive())
	}
}

func TestGetHistoricalMarketCap(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetHistoricalMarketCap(ctx, &model.GetHistoricalMarketCapParams{
		Symbol: "AAPL",
		Since:  types.Date("2025-01-02"),
		Until:  types.Date("2025-01-10"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, mc := range res {
		assert.Equal(t, "AAPL", mc.Symbol)
		assert.NotEmpty(t, string(mc.Date))
		assert.True(t, mc.Value.IsPositive(), "market cap should be positive")
	}
}

func TestGetBulkPriceEOD(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetBulkPriceEOD(ctx, &model.GetBulkPriceEODParams{
		Date: types.Date("2025-01-02"),
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, p := range res[:5] {
		assert.NotEmpty(t, p.Symbol)
		assert.NotEmpty(t, string(p.Date))
		assert.True(t, p.Open.IsPositive())
	}
}

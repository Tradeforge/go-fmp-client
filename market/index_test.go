package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIndexList(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetIndexList(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, e := range res[:5] {
		assert.NotEmpty(t, e.Symbol)
		assert.NotEmpty(t, e.Name)
		assert.NotEmpty(t, e.Exchange)
		assert.NotEmpty(t, e.Currency)
	}
}

func TestBatchGetIndexQuotes(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.BatchGetIndexQuotes(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, q := range res[:5] {
		assert.NotEmpty(t, q.Symbol)
		assert.NotEmpty(t, q.Name)
		assert.NotZero(t, q.Timestamp)
	}
}

func TestBatchGetIndexShortQuotes(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.BatchGetIndexShortQuotes(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, q := range res[:5] {
		assert.NotEmpty(t, q.Symbol)
	}
}

func TestGetHistoricalSP500Constituents(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetHistoricalSP500Constituents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, c := range res[:5] {
		assert.NotEmpty(t, c.Symbol)
		assert.NotEmpty(t, string(c.Date))
		assert.NotEmpty(t, c.AddedSecurity)
	}
}

func TestGetHistoricalNasdaqConstituents(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetHistoricalNasdaqConstituents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, c := range res[:5] {
		assert.NotEmpty(t, c.Symbol)
		assert.NotEmpty(t, string(c.Date))
	}
}

func TestGetHistoricalDowJonesConstituents(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetHistoricalDowJonesConstituents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, c := range res[:5] {
		assert.NotEmpty(t, c.Symbol)
		assert.NotEmpty(t, string(c.Date))
	}
}

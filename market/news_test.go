package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
)

func TestGetFMPArticles(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetFMPArticles(ctx, model.GetFMPArticlesParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title, "article title should not be empty")
		assert.NotEmpty(t, a.Link, "article link should not be empty")
		assert.NotEmpty(t, string(a.Date), "article date should not be empty")
	}
}

func TestGetLatestGeneralNews(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetLatestGeneralNews(ctx, model.GetNewsParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title, "news title should not be empty")
		assert.NotEmpty(t, a.URL, "news URL should not be empty")
		assert.NotEmpty(t, string(a.PublishedDate), "news published date should not be empty")
	}
}

func TestGetPressReleases(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetPressReleases(ctx, model.GetNewsParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, string(a.PublishedDate))
	}
}

func TestGetLatestStockNews(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetLatestStockNews(ctx, model.GetLatestNewsParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, string(a.PublishedDate))
		assert.NotEmpty(t, a.Symbol, "stock news should have a symbol")
	}
}

func TestGetStockNews(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetStockNews(ctx, model.GetNewsParams{Symbols: "AAPL", Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, string(a.PublishedDate))
		assert.Contains(t, a.Symbol, "AAPL")
	}
}

func TestGetLatestCryptoNews(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetLatestCryptoNews(ctx, model.GetLatestNewsParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, string(a.PublishedDate))
	}
}

func TestGetCryptoNews(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetCryptoNews(ctx, model.GetNewsParams{Symbols: "BTCUSD", Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, string(a.PublishedDate))
	}
}

func TestGetLatestForexNews(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetLatestForexNews(ctx, model.GetLatestNewsParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, string(a.PublishedDate))
	}
}

func TestGetForexNews(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := uint(5)
	res, err := client.GetForexNews(ctx, model.GetNewsParams{Symbols: "EURUSD", Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res {
		assert.NotEmpty(t, a.Title)
		assert.NotEmpty(t, string(a.PublishedDate))
	}
}

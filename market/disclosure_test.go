package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
)

func TestGetHouseFinancialDisclosures(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := 10
	res, err := client.GetHouseFinancialDisclosures(ctx, model.GetHouseFinancialDisclosuresParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, d := range res {
		assert.NotEmpty(t, d.FirstName, "first name should not be empty")
		assert.NotEmpty(t, d.LastName, "last name should not be empty")
		assert.NotEmpty(t, string(d.DisclosureDate), "disclosure date should not be empty")
		assert.NotEmpty(t, d.AssetDescription, "asset description should not be empty")
		assert.True(t, d.Type.IsValid(), "disclosure type should be valid: %s", d.Type)
		assert.True(t, d.Amount.Min.IsPositive() || d.Amount.Min.IsZero(), "amount min should be non-negative")
	}
}

func TestGetSenateFinancialDisclosures(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	limit := 10
	res, err := client.GetSenateFinancialDisclosures(ctx, model.GetSenateFinancialDisclosuresParams{Limit: &limit})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, d := range res {
		assert.NotEmpty(t, d.FirstName, "first name should not be empty")
		assert.NotEmpty(t, d.LastName, "last name should not be empty")
		assert.NotEmpty(t, string(d.DisclosureDate), "disclosure date should not be empty")
		assert.True(t, d.Type.IsValid(), "disclosure type should be valid: %s", d.Type)
	}
}

package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
)

func TestGetAdvancedDCF(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetAdvancedDCF(ctx, &model.GetAdvancedDCFParams{Symbol: "AAPL"})
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, "AAPL", res.Symbol)
	assert.True(t, res.Revenue.IsPositive(), "revenue should be positive")
	assert.True(t, res.EBITDA.IsPositive(), "EBITDA should be positive")
	assert.True(t, res.Price.IsPositive(), "price should be positive")
	assert.True(t, res.WACC.IsPositive(), "WACC should be positive")
	assert.True(t, res.EnterpriseValue.IsPositive(), "enterprise value should be positive")
	assert.True(t, res.EquityValuePerShare.IsPositive(), "equity value per share should be positive")
	assert.True(t, res.DilutedSharesOutstanding.IsPositive(), "diluted shares should be positive")
}

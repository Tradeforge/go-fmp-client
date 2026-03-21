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
	require.NotEmpty(t, res)

	assert.Equal(t, "AAPL", res[0].Symbol)
	assert.True(t, res[0].Revenue.IsPositive(), "revenue should be positive")
	assert.True(t, res[0].EBITDA.IsPositive(), "EBITDA should be positive")
	assert.True(t, res[0].Price.IsPositive(), "price should be positive")
	assert.True(t, res[0].WACC.IsPositive(), "WACC should be positive")
	assert.True(t, res[0].EnterpriseValue.IsPositive(), "enterprise value should be positive")
	assert.True(t, res[0].EquityValuePerShare.IsPositive(), "equity value per share should be positive")
	assert.True(t, res[0].DilutedSharesOutstanding.IsPositive(), "diluted shares should be positive")
}

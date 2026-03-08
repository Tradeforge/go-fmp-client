package market

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/model"
)

func TestGetCompanyProfile(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetCompanyProfile(ctx, &model.GetCompanyProfileParams{Symbol: "AAPL"})
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, "AAPL", res.Symbol)
	assert.NotEmpty(t, res.Exchange)
	assert.NotNil(t, res.Price)
	assert.True(t, res.Price.IsPositive())
	assert.NotNil(t, res.CompanyName)
	assert.NotEmpty(t, *res.CompanyName)
	assert.NotNil(t, res.Currency)
	assert.NotNil(t, res.Sector)
	assert.NotNil(t, res.Industry)
	assert.NotNil(t, res.Country)
}

func TestBulkGetCompanyProfile(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.BulkGetCompanyProfile(ctx, &model.BulkGetCompanyProfilesParams{Part: 0})
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, p := range res[:5] {
		assert.NotEmpty(t, p.Symbol)
	}
}

func TestGetFinancialKeyMetricsTTM(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetFinancialKeyMetricsTTM(ctx, &model.GetFinancialKeyMetricsTTMParams{Symbol: "AAPL"})
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, "AAPL", res.Symbol)
	assert.True(t, res.MarketCap.IsPositive(), "market cap should be positive")
	assert.True(t, res.EnterpriseValue.IsPositive(), "enterprise value should be positive")
	assert.False(t, res.ReturnOnEquity.IsZero(), "ROE should not be zero for AAPL")
}

func TestGetFinancialRatiosTTM(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetFinancialRatiosTTM(ctx, &model.GetFinancialRatiosTTMParams{Symbol: "AAPL"})
	require.NoError(t, err)
	require.NotNil(t, res)

	assert.Equal(t, "AAPL", res.Symbol)
	assert.True(t, res.GrossProfitMarginTTM.IsPositive(), "gross profit margin should be positive")
	assert.True(t, res.PriceToEarningsRatioTTM.IsPositive(), "P/E ratio should be positive")
}

func TestGetIncomeStatements(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetIncomeStatements(ctx, &model.GetIncomeStatementsParams{
		Symbol: "AAPL",
		Limit:  4,
		Period: model.FinancialPeriodAnnual,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.LessOrEqual(t, len(res), 4)

	for _, stmt := range res {
		assert.Equal(t, "AAPL", stmt.Symbol)
		assert.NotEmpty(t, string(stmt.Date))
		assert.NotEmpty(t, stmt.ReportedCurrency)
		assert.True(t, stmt.Revenue.IsPositive(), "revenue should be positive")
		assert.True(t, stmt.GrossProfit.IsPositive(), "gross profit should be positive")
		assert.True(t, stmt.NetIncome.IsPositive(), "net income should be positive for AAPL")
	}
}

func TestGetBalanceSheets(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetBalanceSheets(ctx, &model.GetBalanceSheetsParams{
		Symbol: "AAPL",
		Limit:  4,
		Period: model.FinancialPeriodAnnual,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.LessOrEqual(t, len(res), 4)

	for _, bs := range res {
		assert.Equal(t, "AAPL", bs.Symbol)
		assert.NotEmpty(t, bs.Date)
		assert.NotEmpty(t, bs.ReportedCurrency)
		assert.True(t, bs.TotalAssets.IsPositive(), "total assets should be positive")
		assert.True(t, bs.TotalLiabilities.IsPositive(), "total liabilities should be positive")
	}
}

func TestGetCashFlowStatements(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetCashFlowStatements(ctx, &model.GetCashFlowStatementsParams{
		Symbol: "AAPL",
		Limit:  4,
		Period: model.FinancialPeriodAnnual,
	})
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.LessOrEqual(t, len(res), 4)

	for _, cf := range res {
		assert.Equal(t, "AAPL", cf.Symbol)
		assert.NotEmpty(t, string(cf.Date))
		assert.NotEmpty(t, cf.ReportedCurrency)
		assert.True(t, cf.OperatingCashFlow.IsPositive(), "operating cash flow should be positive for AAPL")
	}
}

func TestGetGainers(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetGainers(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, g := range res[:5] {
		assert.NotEmpty(t, g.Symbol)
		assert.NotEmpty(t, g.Name)
		assert.True(t, g.Change.IsPositive(), "gainer change should be positive")
		assert.True(t, g.ChangePercent.IsPositive(), "gainer change percent should be positive")
	}
}

func TestGetLosers(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetLosers(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, l := range res[:5] {
		assert.NotEmpty(t, l.Symbol)
		assert.NotEmpty(t, l.Name)
		assert.True(t, l.Change.IsNegative(), "loser change should be negative")
		assert.True(t, l.ChangePercent.IsNegative(), "loser change percent should be negative")
	}
}

func TestGetMostActiveTickers(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetMostActiveTickers(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, a := range res[:5] {
		assert.NotEmpty(t, a.Symbol)
		assert.NotEmpty(t, a.Name)
		assert.True(t, a.Price.IsPositive(), "price should be positive")
	}
}

func TestGetSP500IndexConstituents(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetSP500IndexConstituents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	assert.GreaterOrEqual(t, len(res), 400, "S&P 500 should have at least 400 constituents")

	for _, c := range res[:5] {
		assert.NotEmpty(t, c.Symbol)
		assert.NotEmpty(t, c.Name)
		assert.NotEmpty(t, c.Sector)
	}
}

func TestGetNasdaqIndexConstituents(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetNasdaqIndexConstituents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, c := range res[:5] {
		assert.NotEmpty(t, c.Symbol)
		assert.NotEmpty(t, c.Name)
	}
}

func TestGetDowJonesIndexConstituents(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetDowJonesIndexConstituents(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	assert.GreaterOrEqual(t, len(res), 25, "Dow Jones should have at least 25 constituents")

	for _, c := range res[:5] {
		assert.NotEmpty(t, c.Symbol)
		assert.NotEmpty(t, c.Name)
	}
}

func TestGetAvailableExchanges(t *testing.T) {
	client := newTestHTTPClient(t)
	ctx := context.Background()

	res, err := client.GetAvailableExchanges(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	for _, e := range res[:5] {
		assert.NotEmpty(t, e.Exchange)
		assert.NotEmpty(t, e.Name)
	}
}

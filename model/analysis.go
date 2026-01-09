package model

import (
	"github.com/shopspring/decimal"
)

type GetAdvancedDCFParams struct {
	Symbol string `query:"symbol,required"`

	// Financial Metrics
	RevenueGrowthPct                  *decimal.Decimal `query:"revenueGrowthPct,omitempty"`
	EBITDAPct                         *decimal.Decimal `query:"ebitdaPct,omitempty"`
	DepreciationAndAmortizationPct    *decimal.Decimal `query:"depreciationAndAmortizationPct,omitempty"`
	CashAndShortTermInvestmentsPct    *decimal.Decimal `query:"cashAndShortTermInvestmentsPct,omitempty"`
	ReceivablesPct                    *decimal.Decimal `query:"receivablesPct,omitempty"`
	InventoriesPct                    *decimal.Decimal `query:"inventoriesPct,omitempty"`
	PayablePct                        *decimal.Decimal `query:"payablePct,omitempty"`
	CapitalExpenditurePct             *decimal.Decimal `query:"capitalExpenditurePct,omitempty"`
	OperatingCashFlowPct              *decimal.Decimal `query:"operatingCashFlowPct,omitempty"`

	// Valuation Assumptions
	TaxRate            *decimal.Decimal `query:"taxRate,omitempty"`
	LongTermGrowthRate *decimal.Decimal `query:"longTermGrowthRate,omitempty"`
	CostOfDebt         *decimal.Decimal `query:"costOfDebt,omitempty"`
	CostOfEquity       *decimal.Decimal `query:"costOfEquity,omitempty"`
	Beta               *decimal.Decimal `query:"beta,omitempty"`
	RiskFreeRate       *decimal.Decimal `query:"riskFreeRate,omitempty"`
}

type GetAdvancedDCFResponse struct {
	Symbol              string          `json:"symbol"`
	Date                string          `json:"date"`
	StockPrice          decimal.Decimal `json:"Stock Price"`
	DCF                 decimal.Decimal `json:"DCF"`
	WACC                decimal.Decimal `json:"WACC"`
	NumberOfShares      decimal.Decimal `json:"numberOfShares"`
	Revenue             decimal.Decimal `json:"revenue"`
	RevenueGrowthRate   decimal.Decimal `json:"revenueGrowthRate"`
	EBITDA              decimal.Decimal `json:"EBITDA"`
	EBITDAMargin        decimal.Decimal `json:"EBITDAMargin"`
	EBIT                decimal.Decimal `json:"EBIT"`
	Depreciation        decimal.Decimal `json:"depreciation"`
	DepreciationPercent decimal.Decimal `json:"depreciationPercent"`
	TotalCash           decimal.Decimal `json:"totalCash"`
	TotalCashPercent    decimal.Decimal `json:"totalCashPercent"`
	Receivables         decimal.Decimal `json:"receivables"`
	ReceivablesPercent  decimal.Decimal `json:"receivablesPercent"`
	Inventories         decimal.Decimal `json:"inventories"`
	InventoriesPercent  decimal.Decimal `json:"inventoriesPercent"`
	Payable             decimal.Decimal `json:"payable"`
	PayablePercent      decimal.Decimal `json:"payablePercent"`
	CapitalExpenditure  decimal.Decimal `json:"capitalExpenditure"`
	CapexPercent        decimal.Decimal `json:"capexPercent"`
	TaxRate             decimal.Decimal `json:"taxRate"`
	LongTermGrowthRate  decimal.Decimal `json:"longTermGrowthRate"`
	RiskFreeRate        decimal.Decimal `json:"riskFreeRate"`
	Beta                decimal.Decimal `json:"beta"`
	CostOfEquity        decimal.Decimal `json:"costOfEquity"`
	CostOfDebt          decimal.Decimal `json:"costOfDebt"`
	FreeCashFlow        decimal.Decimal `json:"freeCashFlow"`
	OperatingCashFlow   decimal.Decimal `json:"operatingCashFlow"`
	OCFPercent          decimal.Decimal `json:"OCFPercent"`
}
package model

import (
    "github.com/shopspring/decimal"
)

type GetAdvancedDCFParams struct {
    Symbol string `query:"symbol,required"`

    // Financial Metrics
    RevenueGrowthPct               *decimal.Decimal `query:"revenueGrowthPct,omitempty"`
    EBITDAPct                      *decimal.Decimal `query:"ebitdaPct,omitempty"`
    EBITPct                        *decimal.Decimal `query:"ebitPct,omitempty"`
    DepreciationAndAmortizationPct *decimal.Decimal `query:"depreciationAndAmortizationPct,omitempty"`
    CashAndShortTermInvestmentsPct *decimal.Decimal `query:"cashAndShortTermInvestmentsPct,omitempty"`
    ReceivablesPct                 *decimal.Decimal `query:"receivablesPct,omitempty"`
    InventoriesPct                 *decimal.Decimal `query:"inventoriesPct,omitempty"`
    PayablePct                     *decimal.Decimal `query:"payablePct,omitempty"`
    CapitalExpenditurePct          *decimal.Decimal `query:"capitalExpenditurePct,omitempty"`
    OperatingCashFlowPct           *decimal.Decimal `query:"operatingCashFlowPct,omitempty"`

    // Valuation Assumptions
    SellingGeneralAndAdministrativeExpensesPct *decimal.Decimal `query:"sellingGeneralAndAdministrativeExpensesPct,omitempty"`
    TaxRate                                    *decimal.Decimal `query:"taxRate,omitempty"`
    LongTermGrowthRate                         *decimal.Decimal `query:"longTermGrowthRate,omitempty"`
    CostOfDebt                                 *decimal.Decimal `query:"costOfDebt,omitempty"`
    CostOfEquity                               *decimal.Decimal `query:"costOfEquity,omitempty"`
    MarketRiskPremium                          *decimal.Decimal `query:"marketRiskPremium,omitempty"`
    Beta                                       *decimal.Decimal `query:"beta,omitempty"`
    RiskFreeRate                               *decimal.Decimal `query:"riskFreeRate,omitempty"`
}

type GetAdvancedDCFResponse struct {
    Year                         string          `json:"year"`
    Symbol                       string          `json:"symbol"`
    Revenue                      decimal.Decimal `json:"revenue"`
    RevenuePercentage            decimal.Decimal `json:"revenuePercentage"`
    EBITDA                       decimal.Decimal `json:"ebitda"`
    EBITDAPercentage             decimal.Decimal `json:"ebitdaPercentage"`
    EBIT                         decimal.Decimal `json:"ebit"`
    EBITPercentage               decimal.Decimal `json:"ebitPercentage"`
    Depreciation                 decimal.Decimal `json:"depreciation"`
    DepreciationPercentage       decimal.Decimal `json:"depreciationPercentage"`
    TotalCash                    decimal.Decimal `json:"totalCash"`
    TotalCashPercentage          decimal.Decimal `json:"totalCashPercentage"`
    Receivables                  decimal.Decimal `json:"receivables"`
    ReceivablesPercentage        decimal.Decimal `json:"receivablesPercentage"`
    Inventories                  decimal.Decimal `json:"inventories"`
    InventoriesPercentage        decimal.Decimal `json:"inventoriesPercentage"`
    Payable                      decimal.Decimal `json:"payable"`
    PayablePercentage            decimal.Decimal `json:"payablePercentage"`
    CapitalExpenditure           decimal.Decimal `json:"capitalExpenditure"`
    CapitalExpenditurePercentage decimal.Decimal `json:"capitalExpenditurePercentage"`
    Price                        decimal.Decimal `json:"price"`
    Beta                         decimal.Decimal `json:"beta"`
    DilutedSharesOutstanding     decimal.Decimal `json:"dilutedSharesOutstanding"`
    CostOfDebt                   decimal.Decimal `json:"costofDebt"`
    TaxRate                      decimal.Decimal `json:"taxRate"`
    AfterTaxCostOfDebt           decimal.Decimal `json:"afterTaxCostOfDebt"`
    RiskFreeRate                 decimal.Decimal `json:"riskFreeRate"`
    MarketRiskPremium            decimal.Decimal `json:"marketRiskPremium"`
    CostOfEquity                 decimal.Decimal `json:"costOfEquity"`
    TotalDebt                    decimal.Decimal `json:"totalDebt"`
    TotalEquity                  decimal.Decimal `json:"totalEquity"`
    TotalCapital                 decimal.Decimal `json:"totalCapital"`
    DebtWeighting                decimal.Decimal `json:"debtWeighting"`
    EquityWeighting              decimal.Decimal `json:"equityWeighting"`
    WACC                         decimal.Decimal `json:"wacc"`
    TaxRateCash                  decimal.Decimal `json:"taxRateCash"`
    EBIAT                        decimal.Decimal `json:"ebiat"`
    UFCF                         decimal.Decimal `json:"ufcf"`
    SumPvUFCF                    decimal.Decimal `json:"sumPvUfcf"`
    LongTermGrowthRate           decimal.Decimal `json:"longTermGrowthRate"`
    TerminalValue                decimal.Decimal `json:"terminalValue"`
    PresentTerminalValue         decimal.Decimal `json:"presentTerminalValue"`
    EnterpriseValue              decimal.Decimal `json:"enterpriseValue"`
    NetDebt                      decimal.Decimal `json:"netDebt"`
    EquityValue                  decimal.Decimal `json:"equityValue"`
    EquityValuePerShare          decimal.Decimal `json:"equityValuePerShare"`
    FreeCashFlowT1               decimal.Decimal `json:"freeCashFlowT1"`
}

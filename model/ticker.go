package model

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetCompanyProfileParams struct {
	Symbol string `query:"symbol,required"`
}

type GetCompanyProfilesParams struct {
	Symbols string `path:"symbols,required"`
}

type BulkGetCompanyProfilesParams struct {
	Part int `query:"part,required" validate:"gte=0"`
}

type IgnoreUnmarshalFailure[T any] struct {
	value *T
}

func (i *IgnoreUnmarshalFailure[T]) Value() *T {
	return i.value
}

func (i *IgnoreUnmarshalFailure[T]) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	if err := json.Unmarshal(b, &i.value); err != nil {
		i.value = nil
	}
	return nil
}

type GetCompanyProfileResponse struct {
	Symbol            string           `json:"symbol"`
	Price             *decimal.Decimal `json:"price,omitempty"`
	Beta              *decimal.Decimal `json:"beta,omitempty"`
	MarketCap         *decimal.Decimal `json:"marketCap,omitempty"`
	LastDividend      *decimal.Decimal `json:"lastDividend,omitempty"`
	Range             *types.Range52w  `json:"range,omitempty"`
	Change            *decimal.Decimal `json:"change,omitempty"`
	ChangePercentage  *decimal.Decimal `json:"changePercentage,omitempty"`
	CompanyName       *string          `json:"companyName,omitempty"`
	Currency          *string          `json:"currency,omitempty"`
	Cik               *string          `json:"cik,omitempty"`
	Isin              *string          `json:"isin,omitempty"`
	Cusip             *string          `json:"cusip,omitempty"`
	ExchangeFullName  string           `json:"exchangeFullName"`
	ExchangeShortName string           `json:"exchangeShortName"`
	Industry          *string          `json:"industry,omitempty"`
	Website           *string          `json:"website,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Ceo               *string          `json:"ceo,omitempty"`
	Sector            *string          `json:"sector,omitempty"`
	Country           *string          `json:"country,omitempty"`
	FullTimeEmployees *decimal.Decimal `json:"fullTimeEmployees,omitempty"`
	Phone             *string          `json:"phone,omitempty"`
	Address           *string          `json:"address,omitempty"`
	City              *string          `json:"city,omitempty"`
	State             *string          `json:"state,omitempty"`
	Zip               *string          `json:"zip,omitempty"`
	DcfDiff           *decimal.Decimal `json:"dcfDiff,omitempty"`
	Dcf               *decimal.Decimal `json:"dcf,omitempty"`
	Image             *string          `json:"image,omitempty"`
	IpoDate           *types.Date      `json:"ipoDate,omitempty"`
	DefaultImage      types.Bool       `json:"defaultImage"`
	IsEtf             types.Bool       `json:"isEtf"`
	IsActivelyTrading types.Bool       `json:"isActivelyTrading"`
	IsAdr             types.Bool       `json:"isAdr"`
	IsFund            types.Bool       `json:"isFund"`
	Volume            *decimal.Decimal `json:"volume,omitempty"`
	VolumeAverage     *decimal.Decimal `json:"averageVolume,omitempty"`
}

type BulkCompanyProfileResponse struct {
	Symbol            string                                                                                 `json:"symbol"`
	Price             types.EmptyOr[decimal.Decimal]                                                         `json:"price,omitempty"`
	MktCap            types.EmptyOr[decimal.Decimal]                                                         `json:"marketCap,omitempty"`
	Beta              types.EmptyOr[decimal.Decimal]                                                         `json:"beta,omitempty"`
	LastDiv           types.EmptyOr[decimal.Decimal]                                                         `json:"lastDividend,omitempty"`
	Range             types.EmptyOr[types.Range52w]                                                          `json:"range,omitempty"`
	Change            types.EmptyOr[decimal.Decimal]                                                         `json:"change,omitempty"`
	ChangePercentage  types.EmptyOr[decimal.Decimal]                                                         `json:"changePercentage,omitempty"`
	Volume            types.EmptyOr[decimal.Decimal]                                                         `json:"volume,omitempty"`
	VolumeAverage     types.EmptyOr[decimal.Decimal]                                                         `json:"averageVolume,omitempty"`
	CompanyName       types.EmptyOr[string]                                                                  `json:"companyName,omitempty"`
	Currency          types.EmptyOr[string]                                                                  `json:"currency,omitempty"`
	Cik               types.EmptyOr[string]                                                                  `json:"cik,omitempty"`
	Isin              types.EmptyOr[string]                                                                  `json:"isin,omitempty"`
	Cusip             types.EmptyOr[string]                                                                  `json:"cusip,omitempty"`
	ExchangeFullName  string                                                                                 `json:"exchangeFullName,omitempty"`
	ExchangeShortName string                                                                                 `json:"exchange,omitempty"`
	Industry          types.EmptyOr[string]                                                                  `json:"industry,omitempty"`
	Website           types.EmptyOr[string]                                                                  `json:"website,omitempty"`
	Description       types.EmptyOr[string]                                                                  `json:"description,omitempty"`
	Ceo               types.EmptyOr[string]                                                                  `json:"ceo,omitempty"`
	Sector            types.EmptyOr[string]                                                                  `json:"sector,omitempty"`
	Country           types.EmptyOr[string]                                                                  `json:"country,omitempty"`
	FullTimeEmployees types.EmptyOr[IgnoreUnmarshalFailure[types.ThousandSeparatedNumeric[decimal.Decimal]]] `json:"fullTimeEmployees,omitempty"`
	Phone             types.EmptyOr[string]                                                                  `json:"phone,omitempty"`
	Address           types.EmptyOr[string]                                                                  `json:"address,omitempty"`
	City              types.EmptyOr[string]                                                                  `json:"city,omitempty"`
	State             types.EmptyOr[string]                                                                  `json:"state,omitempty"`
	Zip               types.EmptyOr[string]                                                                  `json:"zip,omitempty"`
	Image             types.EmptyOr[string]                                                                  `json:"image,omitempty"`
	IpoDate           types.EmptyOr[types.Date]                                                              `json:"ipoDate,omitempty"`
	DefaultImage      types.Bool                                                                             `json:"defaultImage"`
	IsEtf             types.Bool                                                                             `json:"isEtf"`
	IsActivelyTrading types.Bool                                                                             `json:"isActivelyTrading"`
	IsAdr             types.Bool                                                                             `json:"isAdr"`
	IsFund            types.Bool                                                                             `json:"isFund"`
}

func ParseCompanyProfileCSVRecord(header []string, record []string) (*BulkCompanyProfileResponse, error) {
	resultMap := make(map[string]any)
	for j, field := range record {
		resultMap[header[j]] = field
	}
	b, err := json.Marshal(resultMap)
	if err != nil {
		return nil, fmt.Errorf("marshaling record: %w", err)
	}
	var profile BulkCompanyProfileResponse
	if err := json.Unmarshal(b, &profile); err != nil {
		return nil, fmt.Errorf("unmarshaling record: %w", err)
	}
	return &profile, nil
}

type GetFinancialKeyMetricsTTMParams struct {
	Symbol string `query:"symbol,required"`
}

type GetFinancialKeyMetricsTTMResponse = *FinancialKeyMetricsTTM

type FinancialKeyMetricsTTM = FinancialKeyMetrics

type FinancialKeyMetrics struct {
	Symbol                                 string          `json:"symbol"`
	MarketCap                              decimal.Decimal `json:"marketCap"`
	EnterpriseValue                        decimal.Decimal `json:"enterpriseValueTTM"`
	EvToSales                              decimal.Decimal `json:"evToSalesTTM"`
	EvToOperatingCashFlow                  decimal.Decimal `json:"evToOperatingCashFlowTTM"`
	EvToFreeCashFlow                       decimal.Decimal `json:"evToFreeCashFlowTTM"`
	EnterpriseValueOverEBITDA              decimal.Decimal `json:"evToEBITDATTM"`
	NetDebtToEBITDA                        decimal.Decimal `json:"netDebtToEBITDATTM"`
	CurrentRatio                           decimal.Decimal `json:"currentRatioTTM"`
	IncomeQuality                          decimal.Decimal `json:"incomeQualityTTM"`
	GrahamNumber                           decimal.Decimal `json:"grahamNumberTTM"`
	GrahamNetNet                           decimal.Decimal `json:"grahamNetNetTTM"`
	TaxBurden                              decimal.Decimal `json:"taxBurdenTTM"`
	InterestBurden                         decimal.Decimal `json:"interestBurdenTTM"`
	WorkingCapital                         decimal.Decimal `json:"workingCapitalTTM"`
	InvestedCapital                        decimal.Decimal `json:"investedCapitalTTM"`
	ReturnOnAssets                         decimal.Decimal `json:"returnOnAssetsTTM"`
	OperatingReturnOnAssets                decimal.Decimal `json:"operatingReturnOnAssetsTTM"`
	ReturnOnTangibleAssets                 decimal.Decimal `json:"returnOnTangibleAssetsTTM"`
	ReturnOnEquity                         decimal.Decimal `json:"returnOnEquityTTM"`
	ReturnOnInvestedCapital                decimal.Decimal `json:"returnOnInvestedCapitalTTM"`
	ReturnOnCapitalEmployed                decimal.Decimal `json:"returnOnCapitalEmployedTTM"`
	EarningsYield                          decimal.Decimal `json:"earningsYieldTTM"`
	FreeCashFlowYield                      decimal.Decimal `json:"freeCashFlowYieldTTM"`
	CapexToOperatingCashFlow               decimal.Decimal `json:"capexToOperatingCashFlowTTM"`
	CapexToDepreciation                    decimal.Decimal `json:"capexToDepreciationTTM"`
	CapexToRevenue                         decimal.Decimal `json:"capexToRevenueTTM"`
	SalesGeneralAndAdministrativeToRevenue decimal.Decimal `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopmentToRevenue        decimal.Decimal `json:"researchAndDevelopementToRevenueTTM"`
	StockBasedCompensationToRevenue        decimal.Decimal `json:"stockBasedCompensationToRevenueTTM"`
	IntangiblesToTotalAssets               decimal.Decimal `json:"intangiblesToTotalAssetsTTM"`
	AverageReceivables                     decimal.Decimal `json:"averageReceivablesTTM"`
	AveragePayables                        decimal.Decimal `json:"averagePayablesTTM"`
	AverageInventory                       decimal.Decimal `json:"averageInventoryTTM"`
	DaysSalesOutstanding                   decimal.Decimal `json:"daysOfSalesOutstandingTTM"`
	DaysPayablesOutstanding                decimal.Decimal `json:"daysOfPayablesOutstandingTTM"`
	DaysOfInventoryOnHand                  decimal.Decimal `json:"daysOfInventoryOutstandingTTM"`
	OperatingCycle                         decimal.Decimal `json:"operatingCycleTTM"`
	CashConversionCycle                    decimal.Decimal `json:"cashConversionCycleTTM"`
	FreeCashFlowToEquity                   decimal.Decimal `json:"freeCashFlowToEquityTTM"`
	FreeCashFlowToFirm                     decimal.Decimal `json:"freeCashFlowToFirmTTM"`
	TangibleAssetValue                     decimal.Decimal `json:"tangibleAssetValueTTM"`
	NetCurrentAssetValue                   decimal.Decimal `json:"netCurrentAssetValueTTM"`
}

type GetFinancialRatiosTTMParams struct {
	Symbol string `query:"symbol,required"`
}

type GetFinancialRatiosTTMResponse = *FinancialRatiosTTM

type FinancialRatiosTTM struct {
	Symbol string `json:"symbol"`

	// Margin ratios
	GrossProfitMarginTTM                decimal.Decimal `json:"grossProfitMarginTTM"`
	EbitMarginTTM                       decimal.Decimal `json:"ebitMarginTTM"`
	EbitdaMarginTTM                     decimal.Decimal `json:"ebitdaMarginTTM"`
	OperatingProfitMarginTTM            decimal.Decimal `json:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM               decimal.Decimal `json:"pretaxProfitMarginTTM"`
	ContinuousOperationsProfitMarginTTM decimal.Decimal `json:"continuousOperationsProfitMarginTTM"`
	NetProfitMarginTTM                  decimal.Decimal `json:"netProfitMarginTTM"`
	BottomLineProfitMarginTTM           decimal.Decimal `json:"bottomLineProfitMarginTTM"`

	// Activity/Turnover ratios
	ReceivablesTurnoverTTM         decimal.Decimal `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM            decimal.Decimal `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM           decimal.Decimal `json:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM          decimal.Decimal `json:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM               decimal.Decimal `json:"assetTurnoverTTM"`
	WorkingCapitalTurnoverRatioTTM decimal.Decimal `json:"workingCapitalTurnoverRatioTTM"`

	// Liquidity ratios
	CurrentRatioTTM  decimal.Decimal `json:"currentRatioTTM"`
	QuickRatioTTM    decimal.Decimal `json:"quickRatioTTM"`
	SolvencyRatioTTM decimal.Decimal `json:"solvencyRatioTTM"`
	CashRatioTTM     decimal.Decimal `json:"cashRatioTTM"`

	// Valuation ratios
	PriceToEarningsRatioTTM              decimal.Decimal `json:"priceToEarningsRatioTTM"`
	PriceToEarningsGrowthRatioTTM        decimal.Decimal `json:"priceToEarningsGrowthRatioTTM"`
	ForwardPriceToEarningsGrowthRatioTTM decimal.Decimal `json:"forwardPriceToEarningsGrowthRatioTTM"`
	PriceToBookRatioTTM                  decimal.Decimal `json:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                 decimal.Decimal `json:"priceToSalesRatioTTM"`
	PriceToFreeCashFlowRatioTTM          decimal.Decimal `json:"priceToFreeCashFlowRatioTTM"`
	PriceToOperatingCashFlowRatioTTM     decimal.Decimal `json:"priceToOperatingCashFlowRatioTTM"`
	PriceToFairValueTTM                  decimal.Decimal `json:"priceToFairValueTTM"`

	// Debt ratios
	DebtToAssetsRatioTTM          decimal.Decimal `json:"debtToAssetsRatioTTM"`
	DebtToEquityRatioTTM          decimal.Decimal `json:"debtToEquityRatioTTM"`
	DebtToCapitalRatioTTM         decimal.Decimal `json:"debtToCapitalRatioTTM"`
	LongTermDebtToCapitalRatioTTM decimal.Decimal `json:"longTermDebtToCapitalRatioTTM"`
	FinancialLeverageRatioTTM     decimal.Decimal `json:"financialLeverageRatioTTM"`
	DebtToMarketCapTTM            decimal.Decimal `json:"debtToMarketCapTTM"`

	// Cash flow ratios
	OperatingCashFlowRatioTTM                  decimal.Decimal `json:"operatingCashFlowRatioTTM"`
	OperatingCashFlowSalesRatioTTM             decimal.Decimal `json:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM      decimal.Decimal `json:"freeCashFlowOperatingCashFlowRatioTTM"`
	DebtServiceCoverageRatioTTM                decimal.Decimal `json:"debtServiceCoverageRatioTTM"`
	InterestCoverageRatioTTM                   decimal.Decimal `json:"interestCoverageRatioTTM"`
	ShortTermOperatingCashFlowCoverageRatioTTM decimal.Decimal `json:"shortTermOperatingCashFlowCoverageRatioTTM"`
	OperatingCashFlowCoverageRatioTTM          decimal.Decimal `json:"operatingCashFlowCoverageRatioTTM"`
	CapitalExpenditureCoverageRatioTTM         decimal.Decimal `json:"capitalExpenditureCoverageRatioTTM"`

	// Dividend ratios
	DividendPaidAndCapexCoverageRatioTTM decimal.Decimal `json:"dividendPaidAndCapexCoverageRatioTTM"`
	DividendPayoutRatioTTM               decimal.Decimal `json:"dividendPayoutRatioTTM"`
	DividendYieldTTM                     decimal.Decimal `json:"dividendYieldTTM"`

	// Enterprise Value and other metrics
	EnterpriseValueTTM         decimal.Decimal `json:"enterpriseValueTTM"`
	EnterpriseValueMultipleTTM decimal.Decimal `json:"enterpriseValueMultipleTTM"`

	// Per share metrics
	RevenuePerShareTTM            decimal.Decimal `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM          decimal.Decimal `json:"netIncomePerShareTTM"`
	InterestDebtPerShareTTM       decimal.Decimal `json:"interestDebtPerShareTTM"`
	CashPerShareTTM               decimal.Decimal `json:"cashPerShareTTM"`
	BookValuePerShareTTM          decimal.Decimal `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM  decimal.Decimal `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM decimal.Decimal `json:"shareholdersEquityPerShareTTM"`
	OperatingCashFlowPerShareTTM  decimal.Decimal `json:"operatingCashFlowPerShareTTM"`
	CapexPerShareTTM              decimal.Decimal `json:"capexPerShareTTM"`
	FreeCashFlowPerShareTTM       decimal.Decimal `json:"freeCashFlowPerShareTTM"`

	// Other ratios
	NetIncomePerEBTTTM  decimal.Decimal `json:"netIncomePerEBTTTM"`
	EbtPerEbitTTM       decimal.Decimal `json:"ebtPerEbitTTM"`
	EffectiveTaxRateTTM decimal.Decimal `json:"effectiveTaxRateTTM"`
}

type GetMostActiveTickersResponse = []PartialTicker

type GetIncomeStatementsParams struct {
	Symbol string          `query:"symbol,required"`
	Limit  int             `query:"limit,omitempty" validate:"gte=1,lte=120"`
	Period FinancialPeriod `query:"period,omitempty" validate:"oneof=Q1 Q2 Q3 Q4 FY annual quarter"`
}

type PartialTicker struct {
	Symbol        string          `json:"symbol"`
	Name          string          `json:"name"`
	Price         decimal.Decimal `json:"price"`
	Change        decimal.Decimal `json:"change"`
	ChangePercent decimal.Decimal `json:"changePercent"`
}

type GetGainersResponse = []PartialTicker

type GetLosersResponse = []PartialTicker

type GetIndexConstituentsResponse = []IndexConstituent

type IndexConstituent struct {
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	CIK            string  `json:"cik"`
	Sector         string  `json:"sector"`
	DateFirstAdded *string `json:"dateFirstAdded"`
}

type GetAvailableExchangesResponse = []Exchange

type Exchange struct {
	Exchange     string  `json:"exchange" validate:"required"`
	Name         string  `json:"name" validate:"required"`
	SymbolSuffix string  `json:"symbolSuffix"`
	CountryName  *string `json:"countryName"`
	CountryCode  *string `json:"countryCode"`
	Delay        *string `json:"delay"`
}

func (e *Exchange) UnmarshalJSON(b []byte) error {
	type Alias Exchange
	var a Alias
	if err := json.Unmarshal(b, &a); err != nil {
		return fmt.Errorf("unmarshalling exchange: %w", err)
	}
	if a.CountryName != nil && *a.CountryName == "" {
		a.CountryName = nil
	}
	if a.CountryCode != nil && *a.CountryCode == "" {
		a.CountryCode = nil
	}
	if a.SymbolSuffix == "N/A" {
		a.SymbolSuffix = ""
	}
	*e = Exchange(a)
	return nil
}

type GetIncomeStatementsResponse = []IncomeStatement

type IncomeStatement struct {
	Date                                    types.Date      `json:"date"`
	Symbol                                  string          `json:"symbol"`
	ReportedCurrency                        string          `json:"reportedCurrency"`
	Cik                                     string          `json:"cik"`
	FilingDate                              types.Date      `json:"filingDate"`
	AcceptedDateTime                        types.DateTime  `json:"acceptedDate"`
	FiscalYear                              string          `json:"fiscalYear"`
	Period                                  FinancialPeriod `json:"period"`
	Revenue                                 decimal.Decimal `json:"revenue"`
	CostOfRevenue                           decimal.Decimal `json:"costOfRevenue"`
	GrossProfit                             decimal.Decimal `json:"grossProfit"`
	ResearchAndDevelopmentExpenses          decimal.Decimal `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        decimal.Decimal `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             decimal.Decimal `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses decimal.Decimal `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           decimal.Decimal `json:"otherExpenses"`
	OperatingExpenses                       decimal.Decimal `json:"operatingExpenses"`
	CostAndExpenses                         decimal.Decimal `json:"costAndExpenses"`
	NetInterestIncome                       decimal.Decimal `json:"netInterestIncome"`
	InterestIncome                          decimal.Decimal `json:"interestIncome"`
	InterestExpense                         decimal.Decimal `json:"interestExpense"`
	DepreciationAndAmortization             decimal.Decimal `json:"depreciationAndAmortization"`
	Ebitda                                  decimal.Decimal `json:"ebitda"`
	Ebit                                    decimal.Decimal `json:"ebit"`
	NonOperatingIncomeExcludingInterest     decimal.Decimal `json:"nonOperatingIncomeExcludingInterest"`
	OperatingIncome                         decimal.Decimal `json:"operatingIncome"`
	TotalOtherIncomeExpensesNet             decimal.Decimal `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         decimal.Decimal `json:"incomeBeforeTax"`
	IncomeTaxExpense                        decimal.Decimal `json:"incomeTaxExpense"`
	NetIncomeFromContinuingOperations       decimal.Decimal `json:"netIncomeFromContinuingOperations"`
	NetIncomeFromDiscontinuedOperations     decimal.Decimal `json:"netIncomeFromDiscontinuedOperations"`
	OtherAdjustmentsToNetIncome             decimal.Decimal `json:"otherAdjustmentsToNetIncome"`
	NetIncome                               decimal.Decimal `json:"netIncome"`
	NetIncomeDeductions                     decimal.Decimal `json:"netIncomeDeductions"`
	BottomLineNetIncome                     decimal.Decimal `json:"bottomLineNetIncome"`
	Eps                                     decimal.Decimal `json:"eps"`
	EpsDiluted                              decimal.Decimal `json:"epsDiluted"`
	WeightedAverageShsOut                   decimal.Decimal `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                decimal.Decimal `json:"weightedAverageShsOutDil"`
}

type GetBalanceSheetsParams struct {
	Symbol string          `query:"symbol,required"`
	Limit  int             `query:"limit,omitempty" validate:"gte=1,lte=120"`
	Period FinancialPeriod `query:"period,omitempty" validate:"oneof=Q1 Q2 Q3 Q4 FY annual quarter"`
}

type GetBalanceSheetsResponse = []BalanceSheet

type BalanceSheet struct {
	Date                                    string          `json:"date"`
	Symbol                                  string          `json:"symbol"`
	ReportedCurrency                        string          `json:"reportedCurrency"`
	Cik                                     string          `json:"cik"`
	FilingDate                              types.Date      `json:"filingDate"`
	AcceptedDateTime                        types.DateTime  `json:"acceptedDate"`
	FiscalYear                              string          `json:"fiscalYear"`
	Period                                  FinancialPeriod `json:"period"`
	CashAndCashEquivalents                  decimal.Decimal `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    decimal.Decimal `json:"shortTermInvestments"`
	CashAndShortTermInvestments             decimal.Decimal `json:"cashAndShortTermInvestments"`
	NetReceivables                          decimal.Decimal `json:"netReceivables"`
	AccountsReceivables                     decimal.Decimal `json:"accountsReceivables"`
	OtherReceivables                        decimal.Decimal `json:"otherReceivables"`
	Inventory                               decimal.Decimal `json:"inventory"`
	Prepaids                                decimal.Decimal `json:"prepaids"`
	OtherCurrentAssets                      decimal.Decimal `json:"otherCurrentAssets"`
	TotalCurrentAssets                      decimal.Decimal `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               decimal.Decimal `json:"propertyPlantEquipmentNet"`
	Goodwill                                decimal.Decimal `json:"goodwill"`
	IntangibleAssets                        decimal.Decimal `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             decimal.Decimal `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     decimal.Decimal `json:"longTermInvestments"`
	TaxAssets                               decimal.Decimal `json:"taxAssets"`
	OtherNonCurrentAssets                   decimal.Decimal `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   decimal.Decimal `json:"totalNonCurrentAssets"`
	OtherAssets                             decimal.Decimal `json:"otherAssets"`
	TotalAssets                             decimal.Decimal `json:"totalAssets"`
	TotalPayables                           decimal.Decimal `json:"totalPayables"`
	AccountPayables                         decimal.Decimal `json:"accountPayables"`
	OtherPayables                           decimal.Decimal `json:"otherPayables"`
	AccruedExpenses                         decimal.Decimal `json:"accruedExpenses"`
	ShortTermDebt                           decimal.Decimal `json:"shortTermDebt"`
	CapitalLeaseObligationsCurrent          decimal.Decimal `json:"capitalLeaseObligationsCurrent"`
	TaxPayables                             decimal.Decimal `json:"taxPayables"`
	DeferredRevenue                         decimal.Decimal `json:"deferredRevenue"`
	OtherCurrentLiabilities                 decimal.Decimal `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 decimal.Decimal `json:"totalCurrentLiabilities"`
	LongTermDebt                            decimal.Decimal `json:"longTermDebt"`
	DeferredRevenueNonCurrent               decimal.Decimal `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        decimal.Decimal `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              decimal.Decimal `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              decimal.Decimal `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        decimal.Decimal `json:"otherLiabilities"`
	CapitalLeaseObligations                 decimal.Decimal `json:"capitalLeaseObligations"`
	TotalLiabilities                        decimal.Decimal `json:"totalLiabilities"`
	TreasuryStock                           decimal.Decimal `json:"treasuryStock"`
	PreferredStock                          decimal.Decimal `json:"preferredStock"`
	CommonStock                             decimal.Decimal `json:"commonStock"`
	RetainedEarnings                        decimal.Decimal `json:"retainedEarnings"`
	AdditionalPaidInCapital                 decimal.Decimal `json:"additionalPaidInCapital"`
	AccumulatedOtherComprehensiveIncomeLoss decimal.Decimal `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OtherTotalStockholdersEquity            decimal.Decimal `json:"otherTotalStockholdersEquity"`
	TotalStockholdersEquity                 decimal.Decimal `json:"totalStockholdersEquity"`
	TotalEquity                             decimal.Decimal `json:"totalEquity"`
	MinorityInterest                        decimal.Decimal `json:"minorityInterest"`
	TotalLiabilitiesAndTotalEquity          decimal.Decimal `json:"totalLiabilitiesAndTotalEquity"`
	TotalInvestments                        decimal.Decimal `json:"totalInvestments"`
	TotalDebt                               decimal.Decimal `json:"totalDebt"`
	NetDebt                                 decimal.Decimal `json:"netDebt"`
}

type GetCashFlowStatementsParams struct {
	Symbol string          `query:"symbol,required"`
	Limit  int             `query:"limit,omitempty" validate:"gte=1,lte=120"`
	Period FinancialPeriod `query:"period,omitempty" validate:"oneof=Q1 Q2 Q3 Q4 FY annual quarter"`
}

type GetCashFlowStatementsResponse = []CashFlowStatement

type CashFlowStatement struct {
	Date                                   types.Date      `json:"date"`
	Symbol                                 string          `json:"symbol"`
	ReportedCurrency                       string          `json:"reportedCurrency"`
	Cik                                    string          `json:"cik"`
	FilingDate                             types.Date      `json:"filingDate"`
	AcceptedDateTime                       types.DateTime  `json:"acceptedDate"`
	FiscalYear                             string          `json:"fiscalYear"`
	Period                                 FinancialPeriod `json:"period"`
	NetIncome                              decimal.Decimal `json:"netIncome"`
	DepreciationAndAmortization            decimal.Decimal `json:"depreciationAndAmortization"`
	DeferredIncomeTax                      decimal.Decimal `json:"deferredIncomeTax"`
	StockBasedCompensation                 decimal.Decimal `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                 decimal.Decimal `json:"changeInWorkingCapital"`
	AccountsReceivables                    decimal.Decimal `json:"accountsReceivables"`
	Inventory                              decimal.Decimal `json:"inventory"`
	AccountsPayables                       decimal.Decimal `json:"accountsPayables"`
	OtherWorkingCapital                    decimal.Decimal `json:"otherWorkingCapital"`
	OtherNonCashItems                      decimal.Decimal `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities   decimal.Decimal `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment decimal.Decimal `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                        decimal.Decimal `json:"acquisitionsNet"`
	PurchasesOfInvestments                 decimal.Decimal `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments           decimal.Decimal `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivities               decimal.Decimal `json:"otherInvestingActivities"`
	NetCashProvidedByInvestingActivities   decimal.Decimal `json:"netCashProvidedByInvestingActivities"`
	NetDebtIssuance                        decimal.Decimal `json:"netDebtIssuance"`
	LongTermNetDebtIssuance                decimal.Decimal `json:"longTermNetDebtIssuance"`
	ShortTermNetDebtIssuance               decimal.Decimal `json:"shortTermNetDebtIssuance"`
	NetStockIssuance                       decimal.Decimal `json:"netStockIssuance"`
	NetCommonStockIssuance                 decimal.Decimal `json:"netCommonStockIssuance"`
	CommonStockIssuance                    decimal.Decimal `json:"commonStockIssuance"`
	CommonStockRepurchased                 decimal.Decimal `json:"commonStockRepurchased"`
	NetPreferredStockIssuance              decimal.Decimal `json:"netPreferredStockIssuance"`
	NetDividendsPaid                       decimal.Decimal `json:"netDividendsPaid"`
	CommonDividendsPaid                    decimal.Decimal `json:"commonDividendsPaid"`
	PreferredDividendsPaid                 decimal.Decimal `json:"preferredDividendsPaid"`
	OtherFinancingActivities               decimal.Decimal `json:"otherFinancingActivities"`
	NetCashProvidedByFinancingActivities   decimal.Decimal `json:"netCashProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash             decimal.Decimal `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                        decimal.Decimal `json:"netChangeInCash"`
	CashAtEndOfPeriod                      decimal.Decimal `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                decimal.Decimal `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                      decimal.Decimal `json:"operatingCashFlow"`
	CapitalExpenditure                     decimal.Decimal `json:"capitalExpenditure"`
	FreeCashFlow                           decimal.Decimal `json:"freeCashFlow"`
	IncomeTaxesPaid                        decimal.Decimal `json:"incomeTaxesPaid"`
	InterestPaid                           decimal.Decimal `json:"interestPaid"`
}

type FinancialPeriod string

const (
	FinancialPeriodQ1      FinancialPeriod = "Q1"
	FinancialPeriodQ2      FinancialPeriod = "Q2"
	FinancialPeriodQ3      FinancialPeriod = "Q3"
	FinancialPeriodQ4      FinancialPeriod = "Q4"
	FinancialPeriodFY      FinancialPeriod = "FY"
	FinancialPeriodAnnual  FinancialPeriod = "annual"
	FinancialPeriodQuarter FinancialPeriod = "quarter"
)

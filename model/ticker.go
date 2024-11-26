package model

import (
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetCompanyProfileParams struct {
	Symbol string `path:"symbol,required"`
}

type BatchGetCompanyProfilesParams struct {
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
	VolAvg            *decimal.Decimal `json:"volAvg,omitempty"`
	MktCap            *decimal.Decimal `json:"mktCap,omitempty"`
	LastDiv           *decimal.Decimal `json:"lastDiv,omitempty"`
	Range             *types.Range52w  `json:"range,omitempty"`
	Changes           *decimal.Decimal `json:"changes,omitempty"`
	CompanyName       *string          `json:"companyName,omitempty"`
	Currency          *string          `json:"currency,omitempty"`
	Cik               *string          `json:"cik,omitempty"`
	Isin              *string          `json:"isin,omitempty"`
	Cusip             *string          `json:"cusip,omitempty"`
	Exchange          string           `json:"exchange"`
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
}

type BulkCompanyProfileResponse struct {
	Symbol            string                                                                                 `json:"symbol"`
	Price             types.EmptyOr[decimal.Decimal]                                                         `json:"price,omitempty"`
	MktCap            types.EmptyOr[decimal.Decimal]                                                         `json:"marketCap,omitempty"`
	LastDiv           types.EmptyOr[decimal.Decimal]                                                         `json:"lastDividend,omitempty"`
	Range             types.EmptyOr[types.Range52w]                                                          `json:"range,omitempty"`
	Beta              types.EmptyOr[decimal.Decimal]                                                         `json:"beta,omitempty"`
	Change            types.EmptyOr[decimal.Decimal]                                                         `json:"change,omitempty"`
	ChangePercentage  types.EmptyOr[decimal.Decimal]                                                         `json:"changePercentage"`
	CompanyName       types.EmptyOr[string]                                                                  `json:"companyName,omitempty"`
	Currency          types.EmptyOr[string]                                                                  `json:"currency,omitempty"`
	Cik               types.EmptyOr[string]                                                                  `json:"cik,omitempty"`
	Isin              types.EmptyOr[string]                                                                  `json:"isin,omitempty"`
	Cusip             types.EmptyOr[string]                                                                  `json:"cusip,omitempty"`
	Exchange          string                                                                                 `json:"exchangeFullName,omitempty"`
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

type ListHistoricalBarsParams struct {
	Timeframe Timeframe  `path:"timeframe,required"`
	Symbol    string     `path:"symbol,required"`
	Since     types.Date `query:"from"`
	Until     types.Date `query:"to"`
	Extended  bool       `query:"extended"`
}

type ListHistoricalBarsResponse = []Bar

type Bar struct {
	Open     decimal.Decimal `json:"open"`
	High     decimal.Decimal `json:"high"`
	Low      decimal.Decimal `json:"low"`
	Close    decimal.Decimal `json:"close"`
	Volume   decimal.Decimal `json:"volume"`
	DateTime types.DateTime  `json:"date"`
}

type ListStockKeyMetricsParams struct {
	Symbol string `path:"symbol,required"`
	Period string `query:"period,omitempty"`
	Limit  int    `query:"limit,omitempty"`
}

type ListTickerKeyMetricsResponse = []StockKeyMetrics

type StockKeyMetrics struct {
	Symbol                                 string          `json:"symbol"`
	RevenuePerShare                        decimal.Decimal `json:"revenuePerShare"`
	NetIncomePerShare                      decimal.Decimal `json:"netIncomePerShare"`
	OperatingCashFlowPerShare              decimal.Decimal `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare                   decimal.Decimal `json:"freeCashFlowPerShare"`
	CashPerShare                           decimal.Decimal `json:"cashPerShare"`
	BookValuePerShare                      decimal.Decimal `json:"bookValuePerShare"`
	TangibleBookValuePerShare              decimal.Decimal `json:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare             decimal.Decimal `json:"shareholdersEquityPerShare"`
	InterestDebtPerShare                   decimal.Decimal `json:"interestDebtPerShare"`
	MarketCap                              decimal.Decimal `json:"marketCap"`
	EnterpriseValue                        decimal.Decimal `json:"enterpriseValue"`
	PeRatio                                decimal.Decimal `json:"peRatio"`
	PriceToSalesRatio                      decimal.Decimal `json:"priceToSalesRatio"`
	Pocfratio                              decimal.Decimal `json:"pocfratio"`
	PfcfRatio                              decimal.Decimal `json:"pfcfRatio"`
	PbRatio                                decimal.Decimal `json:"pbRatio"`
	PtbRatio                               decimal.Decimal `json:"ptbRatio"`
	EvToSales                              decimal.Decimal `json:"evToSales"`
	EnterpriseValueOverEBITDA              decimal.Decimal `json:"enterpriseValueOverEBITDA"`
	EvToOperatingCashFlow                  decimal.Decimal `json:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       decimal.Decimal `json:"evToFreeCashFlow"`
	EarningsYield                          decimal.Decimal `json:"earningsYield"`
	FreeCashFlowYield                      decimal.Decimal `json:"freeCashFlowYield"`
	DebtToEquity                           decimal.Decimal `json:"debtToEquity"`
	DebtToAssets                           decimal.Decimal `json:"debtToAssets"`
	NetDebtToEBITDA                        decimal.Decimal `json:"netDebtToEBITDA"`
	CurrentRatio                           decimal.Decimal `json:"currentRatio"`
	InterestCoverage                       decimal.Decimal `json:"interestCoverage"`
	IncomeQuality                          decimal.Decimal `json:"incomeQuality"`
	DividendYield                          decimal.Decimal `json:"dividendYield"`
	PayoutRatio                            decimal.Decimal `json:"payoutRatio"`
	SalesGeneralAndAdministrativeToRevenue decimal.Decimal `json:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDdevelopementToRevenue      decimal.Decimal `json:"researchAndDdevelopementToRevenue"`
	IntangiblesToTotalAssets               decimal.Decimal `json:"intangiblesToTotalAssets"`
	CapexToOperatingCashFlow               decimal.Decimal `json:"capexToOperatingCashFlow"`
	CapexToRevenue                         decimal.Decimal `json:"capexToRevenue"`
	CapexToDepreciation                    decimal.Decimal `json:"capexToDepreciation"`
	StockBasedCompensationToRevenue        decimal.Decimal `json:"stockBasedCompensationToRevenue"`
	GrahamNumber                           decimal.Decimal `json:"grahamNumber"`
	Roic                                   decimal.Decimal `json:"roic"`
	ReturnOnTangibleAssets                 decimal.Decimal `json:"returnOnTangibleAssets"`
	GrahamNetNet                           decimal.Decimal `json:"grahamNetNet"`
	WorkingCapital                         decimal.Decimal `json:"workingCapital"`
	TangibleAssetValue                     decimal.Decimal `json:"tangibleAssetValue"`
	NetCurrentAssetValue                   decimal.Decimal `json:"netCurrentAssetValue"`
	InvestedCapital                        decimal.Decimal `json:"investedCapital"`
	AverageReceivables                     decimal.Decimal `json:"averageReceivables"`
	AveragePayables                        decimal.Decimal `json:"averagePayables"`
	AverageInventory                       decimal.Decimal `json:"averageInventory"`
	DaysSalesOutstanding                   decimal.Decimal `json:"daysSalesOutstanding"`
	DaysPayablesOutstanding                decimal.Decimal `json:"daysPayablesOutstanding"`
	DaysOfInventoryOnHand                  decimal.Decimal `json:"daysOfInventoryOnHand"`
	ReceivablesTurnover                    decimal.Decimal `json:"receivablesTurnover"`
	PayablesTurnover                       decimal.Decimal `json:"payablesTurnover"`
	InventoryTurnover                      decimal.Decimal `json:"inventoryTurnover"`
	Roe                                    decimal.Decimal `json:"roe"`
	CapexPerShare                          decimal.Decimal `json:"capexPerShare"`
}

type ListStockRatiosParams struct {
	Symbol string `path:"symbol,required"`
	Period string `query:"period,omitempty"`
	Limit  int    `query:"limit,omitempty"`
}

type ListTickerRatiosResponse = []StockRatios

type StockRatios struct {
	Symbol string `json:"symbol"`

	CashRatio       decimal.Decimal `json:"cashRatio"`
	DebtRatio       decimal.Decimal `json:"debtRatio"`
	DebtEquityRatio decimal.Decimal `json:"debtEquityRatio"`

	GrossProfitMargin     decimal.Decimal `json:"grossProfitMargin"`
	OperatingProfitMargin decimal.Decimal `json:"operatingProfitMargin"`
	PretaxProfitMargin    decimal.Decimal `json:"pretaxProfitMargin"`
	NetProfitMargin       decimal.Decimal `json:"netProfitMargin"`

	EffectiveTaxRate             decimal.Decimal `json:"effectiveTaxRate"`
	ReturnOnAssets               decimal.Decimal `json:"returnOnAssets"`
	ReturnOnEquity               decimal.Decimal `json:"returnOnEquity"`
	ReturnOnCapitalEmployed      decimal.Decimal `json:"returnOnCapitalEmployed"`
	NetIncomePerEBT              decimal.Decimal `json:"netIncomePerEBT"`
	EbtPerEbit                   decimal.Decimal `json:"ebtPerEbit"`
	EbitPerRevenue               decimal.Decimal `json:"ebitPerRevenue"`
	LongTermDebtToCapitalization decimal.Decimal `json:"longTermDebtToCapitalization"`
	TotalDebtToCapitalization    decimal.Decimal `json:"totalDebtToCapitalization"`
	InterestCoverage             decimal.Decimal `json:"interestCoverage"`
	CashFlowToDebtRatio          decimal.Decimal `json:"cashFlowToDebtRatio"`
	CompanyEquityMultiplier      decimal.Decimal `json:"companyEquityMultiplier"`

	ReceivablesTurnover decimal.Decimal `json:"receivablesTurnover"`
	PayablesTurnover    decimal.Decimal `json:"payablesTurnover"`
	InventoryTurnover   decimal.Decimal `json:"inventoryTurnover"`
	FixedAssetTurnover  decimal.Decimal `json:"fixedAssetTurnover"`
	AssetTurnover       decimal.Decimal `json:"assetTurnover"`

	OperatingCashFlowPerShare          decimal.Decimal `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare               decimal.Decimal `json:"freeCashFlowPerShare"`
	CashPerShare                       decimal.Decimal `json:"cashPerShare"`
	OperatingCashFlowSalesRatio        decimal.Decimal `json:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio decimal.Decimal `json:"freeCashFlowOperatingCashFlowRatio"`
	CashFlowCoverageRatios             decimal.Decimal `json:"cashFlowCoverageRatios"`
	ShortTermCoverageRatios            decimal.Decimal `json:"shortTermCoverageRatios"`
	CapitalExpenditureCoverageRatio    decimal.Decimal `json:"capitalExpenditureCoverageRatio"`

	PriceBookValueRatio            decimal.Decimal `json:"priceBookValueRatio"`
	PriceToBookRatio               decimal.Decimal `json:"priceToBookRatio"`
	PriceToSalesRatio              decimal.Decimal `json:"priceToSalesRatio"`
	PriceEarningsRatio             decimal.Decimal `json:"priceEarningsRatio"`
	PriceToFreeCashFlowsRatio      decimal.Decimal `json:"priceToFreeCashFlowsRatio"`
	PriceToOperatingCashFlowsRatio decimal.Decimal `json:"priceToOperatingCashFlowsRatio"`
	PriceCashFlowRatio             decimal.Decimal `json:"priceCashFlowRatio"`
	PriceEarningsToGrowthRatio     decimal.Decimal `json:"priceEarningsToGrowthRatio"`
	PriceSalesRatio                decimal.Decimal `json:"priceSalesRatio"`

	PayoutRatio                       decimal.Decimal `json:"payoutRatio"`
	DividendPaidAndCapexCoverageRatio decimal.Decimal `json:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio               decimal.Decimal `json:"dividendPayoutRatio"`
	DividendYield                     decimal.Decimal `json:"dividendYield"`

	EnterpriseValueMultiple decimal.Decimal `json:"enterpriseValueMultiple"`
}

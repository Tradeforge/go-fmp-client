package model

import (
	"encoding/json"
	"strings"

	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetHouseFinancialDisclosuresParams struct {
	Limit *int `query:"limit" validate:"omitempty,min=1,max=250"`
	Page  *int `query:"page"  validate:"omitempty,min=0,max=100"`
}

type GetHouseFinancialDisclosuresResponse []FinancialDisclosure

type GetSenateFinancialDisclosuresParams struct {
	Limit *int `query:"limit" validate:"omitempty,min=1,max=250"`
	Page  *int `query:"page"  validate:"omitempty,min=0,max=100"`
}

type GetSenateFinancialDisclosuresResponse []FinancialDisclosure

type FinancialDisclosure struct {
	Symbol           string                         `json:"symbol"`
	Type             FinancialDisclosureType        `json:"type"`
	DisclosureDate   types.Date                     `json:"disclosureDate"`
	TransactionDate  types.Date                     `json:"transactionDate"`
	FirstName        string                         `json:"firstName"`
	LastName         string                         `json:"lastName"`
	Office           types.EmptyOr[string]          `json:"office"`
	District         types.EmptyOr[string]          `json:"district"`
	Owner            types.EmptyOr[string]          `json:"owner"`
	AssetDescription string                         `json:"assetDescription"`
	AssetType        string                         `json:"assetType"`
	Amount           FinancialDisclosureRangeAmount `json:"amount"`
	Comment          string                         `json:"comment"`
	Link             string                         `json:"link"`
}

type FinancialDisclosureType string

const (
	FinancialDisclosureTypePurchase FinancialDisclosureType = "Purchase"
	FinancialDisclosureTypeSale     FinancialDisclosureType = "Sale"
)

type FinancialDisclosureRangeAmount struct {
	Range string
	Min   decimal.Decimal
	Max   decimal.Decimal
}

func (r FinancialDisclosureRangeAmount) String() string {
	return r.Range
}

func (r FinancialDisclosureRangeAmount) MarshalJSON() ([]byte, error) {
	if r.Range == "" {
		return nil, nil
	}
	return json.Marshal(r.Range)
}

func (r *FinancialDisclosureRangeAmount) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	r.Range = strings.Replace(s, " ", "", -1) // Remove spaces

	parts := []rune(s)
	sepIndex := -1
	for i, ch := range parts {
		if ch == '-' {
			sepIndex = i
			break
		}
	}
	if sepIndex == -1 {
		return nil // or return an error if preferred
	}
	minStr := string(parts[:sepIndex])
	maxStr := string(parts[sepIndex+1:])
	minDec, err := decimal.NewFromString(minStr)
	if err != nil {
		return err
	}
	maxDec, err := decimal.NewFromString(maxStr)
	if err != nil {
		return err
	}
	r.Min = minDec
	r.Max = maxDec
	return nil
}

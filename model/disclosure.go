package model

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"

	"go.tradeforge.dev/fmp/pkg/types"
)

type GetHouseFinancialDisclosuresParams struct {
	Limit *int  `query:"limit" validate:"omitempty,min=1,max=250"`
	Page  *uint `query:"page"  validate:"omitempty,min=0,max=100"`
}

type GetHouseFinancialDisclosuresResponse []FinancialDisclosure

type GetSenateFinancialDisclosuresParams struct {
	Limit *int  `query:"limit" validate:"omitempty,min=1,max=250"`
	Page  *uint `query:"page"  validate:"omitempty,min=0,max=100"`
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
	Comment          types.EmptyOr[string]          `json:"comment"`
	Link             string                         `json:"link"`
}

type FinancialDisclosureType string

const (
	FinancialDisclosureTypePurchase    FinancialDisclosureType = "Purchase"
	FinancialDisclosureTypeSale        FinancialDisclosureType = "Sale"
	FinancialDisclosureTypeFullSale    FinancialDisclosureType = "Sale (Full)"
	FinancialDisclosureTypePartialSale FinancialDisclosureType = "Sale (Partial)"
)

func (fdt FinancialDisclosureType) IsValid() bool {
	switch fdt {
	case FinancialDisclosureTypePurchase, FinancialDisclosureTypeSale, FinancialDisclosureTypeFullSale, FinancialDisclosureTypePartialSale:
		return true
	}
	return false
}

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
	r.Range = strings.ReplaceAll(s, " ", "")       // Remove spaces
	r.Range = strings.ReplaceAll(r.Range, "$", "") // Remove dollar signs, assume all amounts are in USD
	r.Range = strings.ReplaceAll(r.Range, ",", "") // Remove commas (thousands separator)

	parts := strings.Split(r.Range, "-")
	if len(parts) == 1 {
		// Handle single value case
		valStr := parts[0]
		valDec, err := decimal.NewFromString(valStr)
		if err != nil {
			return err
		}
		r.Min = valDec
		r.Max = valDec
		return nil
	}
	if len(parts) != 2 {
		return fmt.Errorf("invalid FinancialDisclosureRangeAmount range: %s", r.Range)
	}
	minStr := parts[0]
	maxStr := parts[1]
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

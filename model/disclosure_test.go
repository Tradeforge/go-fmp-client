package model

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed fixtures/financial_disclosures.json
var rawFinancialDisclosures string

func TestFinancialDisclosure_UnmarshalJSON(t *testing.T) {
	var disclosures []FinancialDisclosure
	err := json.Unmarshal([]byte(rawFinancialDisclosures), &disclosures)
	require.NoError(t, err)

	for _, d := range disclosures {
		assert.True(t, d.Type.IsValid(), "invalid FinancialDisclosureType: %s", d.Type)
	}
}

func TestFinancialDisclosureRangeAmount_NonNumeric(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{name: "owner prefix", input: `"Spouse/DC Over $1,000,000"`},
		{name: "text only", input: `"Unknown"`},
		{name: "mixed", input: `"$1,001 - Spouse/DCOver1000000"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r FinancialDisclosureRangeAmount
			err := r.UnmarshalJSON([]byte(tt.input))
			assert.NoError(t, err)
		})
	}
}

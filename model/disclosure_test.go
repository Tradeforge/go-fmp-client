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

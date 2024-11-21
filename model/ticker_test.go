package model

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.tradeforge.dev/fmp/pkg/types"
)

func Test_parseCompanyProfileCSVRecord(t *testing.T) {
	type args struct {
		header []string
		record []string
	}
	tests := []struct {
		name    string
		args    args
		testFn  func(response *BulkCompanyProfileResponse)
		wantErr bool
	}{
		{
			name: "success:empty-employee-count",
			args: args{
				header: []string{"Symbol", "FullTimeEmployees"},
				record: []string{"AAPL", ""},
			},
			testFn: func(profile *BulkCompanyProfileResponse) {
				assert.Equal(t, "AAPL", profile.Symbol)
				assert.Nil(t, profile.FullTimeEmployees.Value())
			},
		},
		{
			name: "success:thousand-separated-employee-count",
			args: args{
				header: []string{"Symbol", "FullTimeEmployees"},
				record: []string{"AAPL", "1,234"},
			},
			testFn: func(profile *BulkCompanyProfileResponse) {
				assert.Equal(t, "AAPL", profile.Symbol)
				assert.Equal(t, decimal.NewFromInt(1234), profile.FullTimeEmployees.Value().Value().Value())
			},
		},
		{
			name: "success:ignore-failure-employee-count",
			args: args{
				header: []string{"Symbol", "FullTimeEmployees"},
				record: []string{"AAPL", "{invalid-value}"},
			},
			testFn: func(profile *BulkCompanyProfileResponse) {
				assert.Equal(t, "AAPL", profile.Symbol)
				assert.Nil(t, profile.FullTimeEmployees.Value().Value())
			},
		},
		{
			name: "success:exponentials-52w-range",
			args: args{
				header: []string{"symbol", "range"},
				record: []string{"AAPL", "1.23e-4-5.67e+8"},
			},
			testFn: func(profile *BulkCompanyProfileResponse) {
				assert.Equal(t, "AAPL", profile.Symbol)
				rangeMin, err := decimal.NewFromString("1.23e-4")
				require.NoError(t, err)
				rangeMax, err := decimal.NewFromString("5.67e+8")
				require.NoError(t, err)
				assert.NotNil(t, profile.Range.Value())
				assert.Equal(t, types.Range52w{Min: rangeMin, Max: rangeMax, Sep: "-"}, *profile.Range.Value())
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCompanyProfileCSVRecord(tt.args.header, tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCompanyProfileCSVRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tt.testFn(got)
		})
	}
}

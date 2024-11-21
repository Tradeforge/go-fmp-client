package types

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestParseRange52w(t *testing.T) {
	type args struct {
		r         string
		separator string
	}
	tests := []struct {
		name    string
		args    args
		want    Range52w
		wantErr bool
	}{
		{
			name: "success:valid-range",
			args: args{
				r:         "1.23-4.56",
				separator: "-",
			},
			want: Range52w{
				Min: decimal.NewFromFloat(1.23),
				Max: decimal.NewFromFloat(4.56),
				Sep: "-",
			},
		},
		{
			name: "success:valid-range",
			args: args{
				r:         "12.3e-1-45.6e-1",
				separator: "-",
			},
			want: Range52w{
				Min: decimal.NewFromFloat(1.23),
				Max: decimal.NewFromFloat(4.56),
				Sep: "-",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRange52w(tt.args.r, tt.args.separator)
			if (err != nil) && tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "ParseRange52w(%v, %v)", tt.args.r, tt.args.separator)
		})
	}
}

package percentage_test

import (
	"testing"

	"github.com/ManuchekhrM/convert-number-to-text/pkg/percentage"
	"github.com/ManuchekhrM/convert-number-to-text/valueObjects"
	"github.com/shopspring/decimal"
)

func Test_ConvertNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		lang    string
		want    string
		wantErr error
	}{
		{
			name:    "DecimalNumberRU",
			input:   decimal.NewFromFloat(42.5),
			lang:    "ru",
			want:    "сорок два целых и пять десятых",
			wantErr: nil,
		},
		{
			name:    "DecimalNumberEN",
			input:   decimal.NewFromFloat(3.14),
			lang:    "en",
			want:    "",
			wantErr: valueObjects.ErrInvalidLanguage,
		},
		{
			name:    "InvalidLanguage",
			input:   decimal.NewFromFloat(7.77),
			lang:    "fr",
			want:    "",
			wantErr: valueObjects.ErrInvalidLanguage,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := percentage.ConvertNumber(tt.input, tt.lang)
			if err != tt.wantErr {
				t.Errorf("ConvertNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
			if result != tt.want {
				t.Errorf("ConvertNumber() = %s, want %s", result, tt.want)
			}
		})
	}
}

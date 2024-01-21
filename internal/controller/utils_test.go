package controller

import (
	"github.com/bogdanpashtet/plata-currency-rates/internal/controller/mocks"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_controller_validateIsoCodes(t *testing.T) {
	type fields struct {
		service       *mocks.MockService
		validIsoCodes map[string]struct{}
		logger        zerolog.Logger
	}
	type args struct {
		isoCodes []string
	}
	tests := []struct {
		name        string
		args        args
		want        []string
		wantErr     bool
		invalidElem string
	}{
		{
			name: "ok test. valid input, validation ok",
			args: args{
				isoCodes: []string{"usD", "Eur"},
			},
			want:    []string{"USD", "EUR"},
			wantErr: false,
		},
		{
			name: "fail test. invalid input, validation error",
			args: args{
				isoCodes: []string{"rub", "Eur"},
			},
			want:        []string{"RUB", "Eur"},
			wantErr:     true,
			invalidElem: "RUB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			defer mc.Finish()
			f := fields{
				service:       mocks.NewMockService(mc),
				validIsoCodes: map[string]struct{}{"EUR": {}, "USD": {}},
				logger:        zerolog.Nop(),
			}

			c := New(f.service, f.validIsoCodes, f.logger)

			res, ok := c.validateIsoCode(&tt.args.isoCodes[0], &tt.args.isoCodes[1])

			if tt.wantErr {
				assert.Equal(t, tt.wantErr, !ok)
				assert.Equal(t, tt.invalidElem, res)
			}

			assert.Equal(t, tt.want, tt.args.isoCodes)
		})
	}
}

func Test_controller_getValidIsoCodesString(t *testing.T) {
	type fields struct {
		service       *mocks.MockService
		validIsoCodes map[string]struct{}
		logger        zerolog.Logger
	}
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "ok test. valid input, validation ok",
			want:    "EUR USD",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			defer mc.Finish()
			f := fields{
				service:       mocks.NewMockService(mc),
				validIsoCodes: map[string]struct{}{"EUR": {}, "USD": {}},
				logger:        zerolog.Nop(),
			}

			c := New(f.service, f.validIsoCodes, f.logger)

			res := c.getValidIsoCodesString()

			assert.Equal(t, tt.want, res)
		})
	}
}

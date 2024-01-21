package service

import (
	"context"
	"errors"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models"
	mockCfg "github.com/bogdanpashtet/plata-currency-rates/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_service_GetRateFromProvider(t *testing.T) {
	type fields struct {
		frankfurterPrv *mockCfg.MockFrankfurterPrv
		postgres       *mockCfg.MockPostgres
		logger         zerolog.Logger
	}
	type args struct {
		ctx     context.Context
		toIso   string
		fromIso string
	}
	tests := []struct {
		name        string
		prepare     func(f *fields)
		args        args
		wantErr     bool
		want        models.UpdateResponse
		wantErrText string
	}{
		{
			name: "ok test. frankfurterPrv ok, db ok",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.frankfurterPrv.EXPECT().
						GetRate(gomock.Any(), "USD", "EUR").
						Return([]byte(`{"amount": 1.0,"base": "EUR","date": "2024-01-19","rates": {"USD": 1.0887}}`), nil),
					f.postgres.EXPECT().AddToQueue(gomock.Any(), gomock.Any()).
						Return(nil),
				)
			},
			args: args{
				ctx:     context.Background(),
				toIso:   "USD",
				fromIso: "EUR",
			},
			wantErr: false,
		},
		{
			name: "ok test. frankfurterPrv err",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.frankfurterPrv.EXPECT().
						GetRate(gomock.Any(), "USD", "EUR").
						Return([]byte(``), errors.New("test-error")),
				)
			},
			args: args{
				ctx:     context.Background(),
				toIso:   "USD",
				fromIso: "EUR",
			},
			wantErr:     true,
			wantErrText: "test-error",
		},
		{
			name: "ok test. frankfurterPrv ok; prv invalid json data",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.frankfurterPrv.EXPECT().
						GetRate(gomock.Any(), "USD", "EUR").
						Return([]byte(`hello world`), nil),
				)
			},
			args: args{
				ctx:     context.Background(),
				toIso:   "USD",
				fromIso: "EUR",
			},
			wantErr:     true,
			wantErrText: "invalid character 'h' looking for beginning of value",
		},
		{
			name: "ok test. frankfurterPrv ok; prv invalid data 'base' empty",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.frankfurterPrv.EXPECT().
						GetRate(gomock.Any(), "USD", "EUR").
						Return([]byte(`{"amount": 1.0,"base": "","date": "2024-01-19","rates": {"USD": 1.0887}}`), nil),
				)
			},
			args: args{
				ctx:     context.Background(),
				toIso:   "USD",
				fromIso: "EUR",
			},
			wantErr:     true,
			wantErrText: "get incorrect value from frankfurter",
		},
		{
			name: "ok test. frankfurterPrv ok; prv invalid data 'rates' empty",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.frankfurterPrv.EXPECT().
						GetRate(gomock.Any(), "USD", "EUR").
						Return([]byte(`{"amount": 1.0,"base": "EUR","date": "2024-01-19","rates": {"MXN": 1.0887}}`), nil),
				)
			},
			args: args{
				ctx:     context.Background(),
				toIso:   "USD",
				fromIso: "EUR",
			},
			wantErr:     true,
			wantErrText: "get incorrect value from frankfurter",
		},
		{
			name: "ok test. frankfurterPrv ok; db err",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.frankfurterPrv.EXPECT().
						GetRate(gomock.Any(), "USD", "EUR").
						Return([]byte(`{"amount": 1.0,"base": "EUR","date": "2024-01-19","rates": {"USD": 1.0887}}`), nil),
					f.postgres.EXPECT().AddToQueue(gomock.Any(), gomock.Any()).
						Return(errors.New("test-error")),
				)
			},
			args: args{
				ctx:     context.Background(),
				toIso:   "USD",
				fromIso: "EUR",
			},
			wantErr:     true,
			wantErrText: "test-error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			defer mc.Finish()
			f := fields{
				frankfurterPrv: mockCfg.NewMockFrankfurterPrv(mc),
				postgres:       mockCfg.NewMockPostgres(mc),
				logger:         zerolog.Nop(),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			s := New(f.frankfurterPrv, f.postgres, f.logger)
			_, err := s.GetRateFromProvider(tt.args.ctx, tt.args.toIso, tt.args.fromIso)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tt.wantErrText)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

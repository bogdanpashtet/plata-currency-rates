package models

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCurrencyRateWithDtDto_FromDto(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	testTime := time.Date(2002, 2, 2, 2, 2, 2, 2, loc)

	tests := []struct {
		name    string
		target  CurrencyRateWithDtDto
		want    CurrencyRateWithDt
		wantErr bool
	}{
		{
			name: "ok test",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want: CurrencyRateWithDt{
				Id:       "id",
				Currency: "EUR",
				Base:     "USD",
				Rate:     0.334,
				UpdateDt: testTime,
			},
			wantErr: false,
		},
		{
			name: "fail test/ null value id",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: false},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want:    CurrencyRateWithDt{},
			wantErr: true,
		},
		{
			name: "fail test/ null value currency",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: false},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want:    CurrencyRateWithDt{},
			wantErr: true,
		},
		{
			name: "fail test/ null value base",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: false},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want:    CurrencyRateWithDt{},
			wantErr: true,
		},
		{
			name: "fail test/ null value rate",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: false},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want:    CurrencyRateWithDt{},
			wantErr: true,
		},
		{
			name: "fail test/ null value date",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: false},
			},
			want:    CurrencyRateWithDt{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.target.FromDto()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, result)
				assert.NoError(t, err)
			}

		})
	}
}

func TestCurrencyRateWithDtDto_FromDtoToLast(t *testing.T) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	testTime := time.Date(2002, 2, 2, 2, 2, 2, 2, loc)

	tests := []struct {
		name    string
		target  CurrencyRateWithDtDto
		want    CurrencyRateLast
		wantErr bool
	}{
		{
			name: "ok test",
			target: CurrencyRateWithDtDto{
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want: CurrencyRateLast{
				Currency: "EUR",
				Base:     "USD",
				Rate:     0.334,
				UpdateDt: testTime,
			},
			wantErr: false,
		},
		{
			name: "fail test/ null value currency",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: false},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want:    CurrencyRateLast{},
			wantErr: true,
		},
		{
			name: "fail test/ null value base",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: false},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want:    CurrencyRateLast{},
			wantErr: true,
		},
		{
			name: "fail test/ null value rate",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: false},
				UpdateDt: sql.NullTime{Time: testTime, Valid: true},
			},
			want:    CurrencyRateLast{},
			wantErr: true,
		},
		{
			name: "fail test/ null value date",
			target: CurrencyRateWithDtDto{
				Id:       sql.NullString{String: "id", Valid: true},
				Currency: sql.NullString{String: "EUR", Valid: true},
				Base:     sql.NullString{String: "USD", Valid: true},
				Rate:     sql.NullFloat64{Float64: 0.334, Valid: true},
				UpdateDt: sql.NullTime{Time: testTime, Valid: false},
			},
			want:    CurrencyRateLast{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.target.FromDtoToLast()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want, result)
				assert.NoError(t, err)
			}

		})
	}
}

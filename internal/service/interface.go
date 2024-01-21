package service

import (
	"context"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models"
)

type FrankfurterPrv interface {
	GetRate(ctx context.Context, toIso, fromIso string) ([]byte, error)
}

type Postgres interface {
	AddToQueue(ctx context.Context, rate models.CurrencyRate) error
	ConfirmQueue(ctx context.Context) (models.CurrencyRateWithDt, error)
	GetById(ctx context.Context, id string) (models.CurrencyRateWithDt, error)
	GetLastRate(ctx context.Context, toIso, fromIso string) (models.CurrencyRateLast, error)
}

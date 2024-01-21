package controller

import (
	"context"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models"
)

type Service interface {
	GetRateFromProvider(ctx context.Context, toIso, fromIso string) (models.UpdateResponse, error)
	GetById(ctx context.Context, id string) (models.CurrencyRateWithDt, error)
	GetLastRate(ctx context.Context, toIso, fromIso string) (models.CurrencyRateLast, error)
}

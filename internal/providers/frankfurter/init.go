package frankfurter

import (
	"github.com/bogdanpashtet/plata-currency-rates/internal/infrastructure/requester"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models/config"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

const (
	prvTimeout = 5 * time.Second
)

type provider struct {
	getRate         requester.Requester
	getCurrencyList requester.Requester
	logger          zerolog.Logger
}

func NewProvider(providerCfg *config.Provider, logger zerolog.Logger) *provider {
	httpClient := http.Client{Timeout: prvTimeout}

	return &provider{
		requester.New(&httpClient, *providerCfg, "GetRate"),
		requester.New(&httpClient, *providerCfg, "GetCurrencyList"),
		logger,
	}
}

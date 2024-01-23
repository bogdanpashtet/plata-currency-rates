package main

import (
	"fmt"
	"github.com/bogdanpashtet/plata-currency-rates/internal/bootstrap"
	"github.com/bogdanpashtet/plata-currency-rates/internal/controller"
	"github.com/bogdanpashtet/plata-currency-rates/internal/infrastructure/tech"
	"github.com/bogdanpashtet/plata-currency-rates/internal/providers/frankfurter"
	"github.com/bogdanpashtet/plata-currency-rates/internal/repository/postgres"
	"github.com/bogdanpashtet/plata-currency-rates/internal/router"
	"github.com/bogdanpashtet/plata-currency-rates/internal/service"
	"net/http"
)

var (
	logger = bootstrap.InitLogger()
	cfg    = bootstrap.InitConfig(logger)

	dbConn = bootstrap.DbConnInit(cfg.Postgres, logger)

	frankfurterPrv = frankfurter.NewProvider(&cfg.FrankfurterClient, logger)
	validIsoCodes  = bootstrap.GetValidIsoCodes(frankfurterPrv, logger)
	db             = postgres.New(dbConn, logger)

	svc = service.New(frankfurterPrv, db, logger)
	ctr = controller.New(svc, validIsoCodes, logger)
)

func init() {
	tech.New().SetAppInfo(cfg.Application.Name, cfg.Application.Version)
	bootstrap.StartSyncRates(cfg.SyncRates, svc, logger)
}

func main() {
	s := http.Server{
		Addr:         cfg.Application.Port,
		Handler:      router.NewRouter(ctr),
		ReadTimeout:  cfg.Application.HttpTimeout,
		WriteTimeout: cfg.Application.HttpTimeout,
	}

	logger.Debug().Msg(fmt.Sprintf("server started on port %s", cfg.Application.Port))
	logger.Fatal().Msg(s.ListenAndServe().Error())
}

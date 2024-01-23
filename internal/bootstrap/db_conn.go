package bootstrap

import (
	"context"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models/config"
	"github.com/jackc/pgx"
	"github.com/rs/zerolog"
	"os"
)

func DbConnInit(cfg config.Postgres, logger zerolog.Logger) *pgx.Conn {
	user := os.Getenv(cfg.User)
	if user == "" {
		logger.Fatal().Msg("set env variable for database user")
	}

	password := os.Getenv(cfg.Password)
	if password == "" {
		logger.Fatal().Msg("set env variable for database password")
	}

	pgConfig := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(cfg.Port),
		Database: cfg.Database,
		User:     user,
		Password: password,
	}

	conn, err := pgx.Connect(pgConfig)
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return nil
	}

	err = conn.Ping(context.Background())
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return nil
	}

	logger.Debug().Msg("database connected successfully")

	return conn
}

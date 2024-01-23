package postgres

import (
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"time"
)

const timeout = 10 * time.Second

type database struct {
	conn   *pgx.Conn
	logger zerolog.Logger
}

func New(conn *pgx.Conn, logger zerolog.Logger) *database {
	return &database{
		conn:   conn,
		logger: logger,
	}
}

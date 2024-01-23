package postgres

import (
	"context"
	"errors"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models"
	"github.com/jackc/pgx/v5"
)

func (db *database) AddToQueue(ctx context.Context, rate models.CurrencyRate) error {
	childCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	_, err := db.conn.Exec(childCtx,
		`SELECT * FROM plata_currency_rates.add_to_queue(_id := $1, _currency := $2, _base := $3, _rate := $4)`,
		rate.Id, rate.Currency, rate.Base, rate.Rate)
	if err != nil {
		db.logger.Error().Msg(err.Error())
		return err
	}

	return err
}

func (db *database) ConfirmQueue(ctx context.Context) (models.CurrencyRateWithDt, error) {
	childCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var currRate models.CurrencyRateDto
	var rate models.CurrencyRateWithDtDto

	tx, err := db.conn.Begin(childCtx)
	if err != nil {
		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateWithDt{}, err
	}

	defer tx.Rollback(childCtx)

	err = tx.QueryRow(childCtx,
		`SELECT * FROM plata_currency_rates.confirm_queue();`).
		Scan(&currRate.Id, &currRate.Currency, &currRate.Base, &currRate.Rate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			db.logger.Warn().Msg("queue is empty")
			return models.CurrencyRateWithDt{}, err
		}

		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateWithDt{}, err
	}

	err = tx.QueryRow(childCtx,
		`SELECT * FROM plata_currency_rates.add_to_rates(_id := $1, _currency := $2, _base := $3, _rate := $4);`,
		currRate.Id, currRate.Currency, currRate.Base, currRate.Rate,
	).Scan(&rate.Id, &rate.Currency, &rate.Base, &rate.Rate, &rate.UpdateDt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			db.logger.Warn().Msg(err.Error())
			return models.CurrencyRateWithDt{}, err
		}

		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateWithDt{}, err
	}

	result, err := rate.FromDto()
	if err != nil {
		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateWithDt{}, err
	}

	err = tx.Commit(childCtx)
	if err != nil {
		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateWithDt{}, err
	}

	return result, err
}

func (db *database) GetById(ctx context.Context, id string) (models.CurrencyRateWithDt, error) {
	childCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var rate models.CurrencyRateWithDtDto

	err := db.conn.QueryRow(childCtx,
		`SELECT * FROM plata_currency_rates.get_by_id(_id := $1);`,
		id).
		Scan(&rate.Id, &rate.Currency, &rate.Base, &rate.Rate, &rate.UpdateDt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			db.logger.Warn().Msg(err.Error())
			return models.CurrencyRateWithDt{}, err
		}

		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateWithDt{}, err
	}

	result, err := rate.FromDto()
	if err != nil {
		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateWithDt{}, err
	}

	return result, err
}

func (db *database) GetLastRate(ctx context.Context, toIso, fromIso string) (models.CurrencyRateLast, error) {
	childCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var rate models.CurrencyRateWithDtDto

	err := db.conn.QueryRow(childCtx,
		`SELECT * FROM plata_currency_rates.get_last_rate(_currency := $1, _base := $2);`,
		toIso, fromIso).
		Scan(&rate.Currency, &rate.Base, &rate.Rate, &rate.UpdateDt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			db.logger.Warn().Msg(err.Error())
			return models.CurrencyRateLast{}, err
		}

		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateLast{}, err
	}

	result, err := rate.FromDtoToLast()
	if err != nil {
		db.logger.Error().Msg(err.Error())
		return models.CurrencyRateLast{}, err
	}

	return result, err
}

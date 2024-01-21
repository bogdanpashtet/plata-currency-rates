package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bogdanpashtet/plata-currency-rates/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

func (svc *service) GetRateFromProvider(ctx context.Context, toIso, fromIso string) (models.UpdateResponse, error) {
	respBody, err := svc.frankfurterPrv.GetRate(ctx, toIso, fromIso)
	if err != nil {
		return models.UpdateResponse{}, err
	}

	var rate map[string]interface{}
	if err_ := json.Unmarshal(respBody, &rate); err_ != nil {
		svc.logger.Error().Msg(err_.Error())
		return models.UpdateResponse{}, err_
	}

	if rate["rates"].(map[string]interface{})[toIso] == nil || rate["base"] != fromIso {
		err_ := errors.New("get incorrect value from frankfurter")
		svc.logger.Error().Msg(err_.Error())
		return models.UpdateResponse{}, err_
	}

	currRate := models.CurrencyRate{
		Id:       uuid.New().String(),
		Currency: toIso,
		Base:     fromIso,
		Rate:     float32(rate["rates"].(map[string]interface{})[toIso].(float64)),
	}

	if err_ := svc.db.AddToQueue(ctx, currRate); err_ != nil {
		return models.UpdateResponse{}, err_
	}

	svc.logger.Info().Msg(fmt.Sprintf("succesfully added to queue: %+v", currRate))

	rateResp := models.UpdateResponse{RateId: currRate.Id}

	return rateResp, err
}

func (svc *service) GetById(ctx context.Context, id string) (models.CurrencyRateWithDt, error) {
	return svc.db.GetById(ctx, id)
}

func (svc *service) GetLastRate(ctx context.Context, toIso, fromIso string) (models.CurrencyRateLast, error) {
	return svc.db.GetLastRate(ctx, toIso, fromIso)
}

func (svc *service) SyncRates() {
	ctx := context.Background()

	rate, err := svc.db.ConfirmQueue(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return
		}

		svc.logger.Error().Msg(err.Error())
		return
	}

	svc.logger.Info().Msg(fmt.Sprintf("rate updated successfully: %+v", rate))
}

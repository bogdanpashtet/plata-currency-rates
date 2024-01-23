package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bogdanpashtet/plata-currency-rates/internal/infrastructure/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"strings"
)

// UpdateRate godoc
// @Summary      	Send signal to update rate
// @Tags         	Methods
// @Param 			rate query string false "currency rate" example(EUR/USD)
// @Success      	200 {object} models.UpdateResponse "success"
// @Failure      	400 "validation error"
// @Failure      	500 "service unavailable"
// @Router       	/ [put]
func (ctr *controller) UpdateRate(w http.ResponseWriter, r *http.Request) {
	currencyRate := r.URL.Query().Get("rate")
	currencies := strings.Split(currencyRate, "/")
	if len(currencies) != 2 {
		err_ := errors.New("parameter doesn't match pattern EUR/USD")
		ctr.logger.Error().Msg(err_.Error())
		response.WriteError(w, http.StatusBadRequest, err_)
		return
	}

	if invalidIso, isInvalid := ctr.validateIsoCode(&currencies[0], &currencies[1]); !isInvalid {
		err_ := fmt.Errorf("uexpected iso code %s. try this one: %s", invalidIso, ctr.getValidIsoCodesString())
		ctr.logger.Error().Msg(fmt.Sprintf("uexpected iso code %s", invalidIso))
		response.WriteError(w, http.StatusBadRequest, err_)
		return
	}

	rateId, err := ctr.service.GetRateFromProvider(r.Context(), currencies[0], currencies[1])
	if err != nil {
		ctr.logger.Error().Msg(err.Error())
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	respBody, err := json.Marshal(rateId)
	if err != nil {
		ctr.logger.Error().Msg(err.Error())
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response.Write(w, respBody)
}

// GetById godoc
// @Summary      	Get currency rate by id
// @Tags         	Methods
// @Param 			id path string false "currency rate update ID" example(ed7f018b-dc91-4940-8d57-4f91cfe5a8bc)
// @Success      	200 {object} models.CurrencyRateWithDt "success"
// @Failure      	400 "validation error"
// @Failure      	500 "service unavailable"
// @Router       	/by-id/{id} [get]
func (ctr *controller) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" {
		err_ := errors.New("set id value")
		ctr.logger.Error().Msg(err_.Error())
		response.WriteError(w, http.StatusBadRequest, err_)
		return
	}

	if err_ := uuid.Validate(id); err_ != nil {
		ctr.logger.Error().Msg(err_.Error())
		response.WriteError(w, http.StatusBadRequest, err_)
		return
	}

	result, err := ctr.service.GetById(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ctr.logger.Error().Msg(err.Error())
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	respBody, err := json.Marshal(result)
	if err != nil {
		ctr.logger.Error().Msg(err.Error())
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response.Write(w, respBody)
}

// GetLastRate godoc
// @Summary      	Get latest currency rate
// @Tags         	Methods
// @Param 			rate query string false "currency rate" example(EUR/USD)
// @Success      	200 {object} models.CurrencyRateLast "success"
// @Failure      	400 "validation error"
// @Failure      	500 "service unavailable"
// @Router       	/last [get]
func (ctr *controller) GetLastRate(w http.ResponseWriter, r *http.Request) {
	currencyRate := r.URL.Query().Get("rate")
	currencies := strings.Split(currencyRate, "/")
	if len(currencies) != 2 {
		err_ := errors.New("parameter doesn't match pattern EUR/USD")
		ctr.logger.Error().Msg(err_.Error())
		response.WriteError(w, http.StatusBadRequest, err_)
		return
	}

	if invalidIso, isInvalid := ctr.validateIsoCode(&currencies[0], &currencies[1]); !isInvalid {
		err_ := fmt.Errorf("uexpected iso code %s. try this one: %s", invalidIso, ctr.getValidIsoCodesString())
		ctr.logger.Error().Msg(fmt.Sprintf("uexpected iso code %s", invalidIso))
		response.WriteError(w, http.StatusBadRequest, err_)
		return
	}

	result, err := ctr.service.GetLastRate(r.Context(), currencies[0], currencies[1])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		ctr.logger.Error().Msg(err.Error())
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	respBody, err := json.Marshal(result)
	if err != nil {
		ctr.logger.Error().Msg(err.Error())
		response.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response.Write(w, respBody)
}

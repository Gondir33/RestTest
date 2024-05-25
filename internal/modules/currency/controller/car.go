package controller

import (
	"goTest/internal/infrastructure/responder"
	"goTest/internal/modules/currency/service"
	"net/http"

	"goTest/internal/infrastructure/godecoder"

	"go.uber.org/zap"
)

type CurrencyHandler interface {
	GetCurrency(w http.ResponseWriter, r *http.Request)
}

type Currency struct {
	Service   service.CurrencyService
	Responder responder.Responder
	Decoder   godecoder.Decoder
	Logger    *zap.Logger
}

func NewCurrencyHandler(Сar service.CurrencyService, respond responder.Responder, Decoder godecoder.Decoder, Logger *zap.Logger) CurrencyHandler

// @Summary	GetCurrencys with filter and pagination
// @Tags		Currency
// @Accept		json
// @Produce	json
// @Param		regNum		query		string	false	"string"
// @Param		mark		query		string	false	"string"
// @Param		model		query		string	false	"string"
// @Param		year		query		int		false	"2002"
// @Param		name		query		string	false	"string"
// @Param		surname		query		string	false	"string"
// @Param		patronymic	query		string	false	"string"
// @Success	200			{object}	[]models.Currency
// @Failure      400
// @Failure      500
// @Router		/Currency [get]
// func (c Currency) GetCurrency(w http.ResponseWriter, r *http.Request) {

// 	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
// 	if err != nil {
// 		c.Logger.Debug("get Currency bad request", zap.Error(err))
// 		c.Responder.ErrorBadRequest(w, err)
// 		return
// 	}
// 	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
// 	if err != nil {
// 		c.Logger.Debug("get Currency bad request", zap.Error(err))
// 		c.Responder.ErrorBadRequest(w, err)
// 		return
// 	}
// 	Currencys, err := c.Service.GetCurrencys(r.Context(), filters, limit, offset)
// 	if err != nil {
// 		c.Logger.Debug("get Currency internal error", zap.Error(err))
// 		c.Responder.ErrorInternal(w, err)
// 		return
// 	}
// 	c.Responder.OutputJSON(w, Currencys)
// }

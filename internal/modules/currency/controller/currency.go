package controller

import (
	"fmt"
	"goTest/internal/infrastructure/responder"
	"goTest/internal/modules/currency/service"
	"net/http"
	"net/url"

	"goTest/internal/infrastructure/godecoder"

	"go.uber.org/zap"
)

type CurrencyHandler interface {
	GetCurrency(w http.ResponseWriter, r *http.Request)
}

type Currency struct {
	service   service.CurrencyService
	responder responder.Responder
	decoder   godecoder.Decoder
	log       *zap.Logger
}

func NewCurrencyHandler(curr service.CurrencyService, respond responder.Responder, Decoder godecoder.Decoder, Logger *zap.Logger) CurrencyHandler {
	return &Currency{
		service:   curr,
		responder: respond,
		decoder:   Decoder,
		log:       Logger,
	}
}

// @Summary	GetCurrencys with filter and pagination
// @Tags		Currency
// @Accept		json
// @Produce	json
// @Param		date 		query		string	false	"string"
// @Param		val 		query		string	true	"string"
// @Success	200			{object}	[]models.Currency
// @Failure      400	string
// @Failure      500	string
// @Router		/Currency [get]
func (c Currency) GetCurrency(w http.ResponseWriter, r *http.Request) {

	date, currencyName, err := c.CheckInput(r.URL)

	if err != nil {
		c.responder.ErrorBadRequest(w, err)
		return
	}

	CurrencyModel, err := c.service.GetCurrency(r.Context(), date, currencyName)
	if err != nil {
		c.responder.ErrorInternal(w, err)
		return
	}

	c.responder.OutputJSON(w, CurrencyModel)
}

func (c Currency) CheckInput(URL *url.URL) (string, string, error) {
	date := URL.Query().Get("date")
	currencyName := URL.Query().Get("val")
	if currencyName == "" {
		return "", "", fmt.Errorf("Need to Name of Currency in query like val=<name of Currency>")
	}
	return date, currencyName, nil
}

package controller

import (
	"fmt"
	"goTest/internal/infrastructure/component"
	"goTest/internal/infrastructure/responder"
	"goTest/internal/modules/currency/service"
	"net/http"
	"net/url"
	"time"

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

func NewCurrencyHandler(curr service.CurrencyService, components *component.Components) CurrencyHandler {
	return &Currency{
		service:   curr,
		responder: components.Responder,
		decoder:   components.Decoder,
		log:       components.Logger,
	}
}

// @Summary	GetCurrencys with date and val
// @Tags		Currency
// @Accept		json
// @Produce	json
// @Param		date	query		string	false	"example: 02/03/2002"
// @Param		val		query		string	true	"example: USD"
// @Success	200		{object}	models.Currency
// @Failure	400		string		string
// @Failure	500		string		string
// @Router		/currency [get]
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
		return "", "", fmt.Errorf("need to Code of Valute in query like val=<Code of Valute>")
	}

	if _, ok := mapOfCode[currencyName]; !ok {
		return "", "", fmt.Errorf("no such valute in base, we have only: %v", mapOfCode)
	}
	if date != "" {
		dateTime, err := time.Parse("02/01/2006", date)
		if err != nil {
			return "", "", fmt.Errorf("not valid date: %w", err)
		}

		if dateTime.Unix() > time.Now().Unix() {
			return "", "", fmt.Errorf("FutureDate")
		}
	}

	if _, ok := mapOfCode[currencyName]; !ok {
		return "", "", fmt.Errorf("not dalid date")
	}

	return date, currencyName, nil
}

var mapOfCode = map[string]struct{}{
	"AUD": struct{}{},
	"AZN": struct{}{},
	"GBP": struct{}{},
	"AMD": struct{}{},
	"BYN": struct{}{},
	"BGN": struct{}{},
	"BRL": struct{}{},
	"HUF": struct{}{},
	"VND": struct{}{},
	"HKD": struct{}{},
	"GEL": struct{}{},
	"DKK": struct{}{},
	"AED": struct{}{},
	"USD": struct{}{},
	"EUR": struct{}{},
	"EGP": struct{}{},
	"INR": struct{}{},
	"IDR": struct{}{},
	"KZT": struct{}{},
	"CAD": struct{}{},
	"QAR": struct{}{},
	"KGS": struct{}{},
	"CNY": struct{}{},
	"MDL": struct{}{},
	"NZD": struct{}{},
	"NOK": struct{}{},
	"PLN": struct{}{},
	"RON": struct{}{},
	"XDR": struct{}{},
	"SGD": struct{}{},
	"TJS": struct{}{},
	"THB": struct{}{},
	"TRY": struct{}{},
	"TMT": struct{}{},
	"UZS": struct{}{},
	"UAH": struct{}{},
	"CZK": struct{}{},
	"SEK": struct{}{},
	"CHF": struct{}{},
	"RSD": struct{}{},
	"ZAR": struct{}{},
	"KRW": struct{}{},
	"JPY": struct{}{},
}

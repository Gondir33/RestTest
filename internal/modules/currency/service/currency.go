package service

import (
	"context"
	"fmt"
	"goTest/internal/infrastructure/component"
	currencyapi "goTest/internal/infrastructure/currencyApi"
	"goTest/internal/infrastructure/godecoder"
	"goTest/internal/models"
	"goTest/internal/modules/currency/storage"

	"go.uber.org/zap"
)

//go:generate mockgen -source=currency.go -destination=mocks/currency_mock.go
type CurrencyService interface {
	GetCurrency(ctx context.Context, date, currencyName string) (models.Currency, error)
}

type Currency struct {
	CurrencyRep storage.CurrencyStorage
	decoder     godecoder.Decoder
	log         *zap.Logger
}

func NewCurrencyService(CurrencyRep storage.CurrencyStorage, components *component.Components) CurrencyService {
	return &Currency{
		CurrencyRep: CurrencyRep,
		decoder:     components.Decoder,
		log:         components.Logger,
	}
}

func (c *Currency) GetCurrency(ctx context.Context, date, currencyName string) (models.Currency, error) {

	select {
	case <-ctx.Done():
		return models.Currency{}, fmt.Errorf("a lot of wait time of api")
	default:
		return currencyapi.GetCurrencyByNameDate(date, currencyName)
	}
}

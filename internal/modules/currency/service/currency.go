package service

import (
	"context"
	"goTest/internal/infrastructure/component"
	"goTest/internal/infrastructure/godecoder"
	"goTest/internal/models"
	"goTest/internal/modules/currency/storage"

	"go.uber.org/zap"
)

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

func (c *Currency) GetCurrency(ctx context.Context, date, currencyName string) (models.Currency, error)

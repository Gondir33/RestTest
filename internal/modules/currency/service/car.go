package service

import (
	"goTest/internal/infrastructure/component"
	"goTest/internal/modules/currency/storage"
)

type CurrencyService interface {
	GetCurrency()
}

type Currency struct {
}

func NewCurrencyService(CarerRep storage.CurrencyStorage, components *component.Components) CurrencyService

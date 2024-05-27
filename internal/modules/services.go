package modules

import (
	"goTest/internal/infrastructure/component"
	cService "goTest/internal/modules/currency/service"
	"goTest/internal/storages"
)

type Services struct {
	Service cService.CurrencyService
}

func NewServices(storages *storages.Storages, components *component.Components) *Services {
	return &Services{
		Service: cService.NewCurrencyService(storages.CurrencyStorage, components),
	}
}

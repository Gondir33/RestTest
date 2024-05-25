package modules

import (
	"goTest/internal/infrastructure/component"
	cHandler "goTest/internal/modules/currency/controller"
)

type Controllers struct {
	Currency cHandler.CurrencyHandler
}

func NewControllers(services *Services, components *component.Components) *Controllers

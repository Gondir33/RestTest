package storages

import (
	cStorage "goTest/internal/modules/currency/storage"
)

type Storages struct {
	cStorage.CurrencyStorage
}

func NewStorages(curr cStorage.CurrencyStorage) *Storages {
	return &Storages{
		CurrencyStorage: curr,
	}
}

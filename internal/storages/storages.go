package storages

import (
	cStorage "goTest/internal/modules/currency/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storages struct {
	cStorage.CurrencyStorage
}

func NewStorages(pool *pgxpool.Pool) *Storages

package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type CurrencyStorage interface {
}
type Currency struct {
	pool *pgxpool.Pool
}

func NewCurrencyStorage(pool *pgxpool.Pool) CurrencyStorage

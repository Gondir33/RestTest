package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type CurrencyStorage interface {
	InsertCurrency(date, name string)
}
type Currency struct {
	pool *pgxpool.Pool
}

func NewCurrencyStorage(pool *pgxpool.Pool) CurrencyStorage {
	return &Currency{
		pool: pool,
	}
}

func (c *Currency) InsertCurrency(date, name string)

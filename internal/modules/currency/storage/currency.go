package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:generate mockgen -source=currency.go -destination=mocks/currency_mock.go
type CurrencyStorage interface {
	InsertCurrency(ctx context.Context, name string, value float64) error
}
type Currency struct {
	pool *pgxpool.Pool
}

func NewCurrencyStorage(pool *pgxpool.Pool) CurrencyStorage {
	return &Currency{
		pool: pool,
	}
}

func (c *Currency) InsertCurrency(ctx context.Context, name string, value float64) error {
	sql := `INSERT INTO currency (name, value, created_at)
	VALUES ($1, $2, $3)`

	_, err := c.pool.Exec(ctx, sql, name, value, time.Now())

	return err
}

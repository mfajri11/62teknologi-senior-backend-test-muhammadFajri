package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Option func(*pgxpool.Config)

func WithMaxPoolSize(n int32) Option {
	return func(c *pgxpool.Config) {
		c.MaxConns = n
	}
}

func WithMinxPoolSize(n int32) Option {
	return func(c *pgxpool.Config) {
		c.MinConns = n
	}
}

func WithMaxConnLifeTime(t time.Duration) Option {
	return func(c *pgxpool.Config) {
		c.MaxConnLifetime = t
	}
}

func WithMaxConnIdleTime(t time.Duration) Option {
	return func(c *pgxpool.Config) {
		c.MaxConnIdleTime = t
	}
}

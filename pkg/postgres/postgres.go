package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLer interface {
	// Close() error
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
}

const defaultAttempts int = 5

func New(connString string, opts ...Option) *pgxpool.Pool {
	var (
		err  error
		conn *pgxpool.Pool
	)
	// check pgxpool implement SQLer
	var _ SQLer = (*pgxpool.Pool)(nil)
	// if options exists create pool with those options values, else use default
	if len(opts) > 0 {
		cfg, errParse := pgxpool.ParseConfig(connString)
		if errParse != nil {
			errParse = fmt.Errorf("postgres.New: error parse config: %w", errParse)
			fmt.Fprint(os.Stderr, errParse.Error())
			os.Exit(1)
		}
		for _, opt := range opts {
			opt(cfg)
		}

		conn, err = pgxpool.NewWithConfig(context.Background(), cfg)
	} else {
		conn, err = pgxpool.New(context.Background(), connString)
	}
	if err != nil {
		err := fmt.Errorf("postgres.New: error connect to postgres: %w", err)
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	for i := defaultAttempts; i > 0; i-- {
		err = conn.Ping(context.Background())
		if err == nil {
			break
		}
		fmt.Fprintf(os.Stdout, "postgres.New: fail to connect to postgres, attempts left %d", i)
	}
	if err != nil {
		err = fmt.Errorf("postgres.New: error ping to postgres: %w", err)
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
	return conn
}

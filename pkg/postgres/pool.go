package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgClient interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error)
}

func NewPool(ctx context.Context, pgCfg pgConfig, logger Logger) (*pgxpool.Pool, error) {
	const op = "postgres.NewPool: %v"

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&pool_max_conns=%d",
		pgCfg.Username,
		pgCfg.Password,
		pgCfg.Host,
		pgCfg.Port,
		pgCfg.Database,
		pgCfg.SSLMode,
		pgCfg.PoolMaxConns,
	)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	err = tryToPingWithAttempts(ctx, pool, pgCfg.MaxAttempts, pgCfg.MaxDelay, logger)
	if err != nil {
		return nil, fmt.Errorf(op, err)
	}

	return pool, nil
}

func tryToPingWithAttempts(
	ctx context.Context,
	pool *pgxpool.Pool,
	attempts int,
	delay time.Duration,
	logger Logger,
) (err error) {
	for i := 0; i < attempts; i++ {
		err = pool.Ping(ctx)
		if err == nil {
			return nil
		}

		msg := fmt.Sprintf("failed to connect to database with error: %v. Trying again", err)
		logger.Error(msg)

		time.Sleep(delay)
	}
	return err
}

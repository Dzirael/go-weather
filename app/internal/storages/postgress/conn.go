package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-weather/app/internal/config"
	"go-weather/app/internal/storages/postgress/sqlc"
	"go-weather/migrations"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
)

func NewPostgres(lc fx.Lifecycle, cfg *config.Config) (*pgxpool.Pool, *sqlc.Queries, error) {
	connectCfg, err := pgxpool.ParseConfig(cfg.DB.PostgresDSN)
	if err != nil {
		return nil, nil, fmt.Errorf("parse postgres dsn: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), connectCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("create pool: %w", err)
	}

	if err := runMigration(cfg.DB.PostgresDSN); err != nil {
		return nil, nil, fmt.Errorf("run migrations: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return pool.Ping(ctx)
		},
		OnStop: func(_ context.Context) error {
			pool.Close()
			return nil
		},
	})

	return pool, sqlc.New(pool), nil
}

func runMigration(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer db.Close()

	goose.SetBaseFS(migrations.Migrations)

	if err := goose.Up(db, "."); err != nil {
		return fmt.Errorf("goose up (embedded): %w", err)
	}

	return nil

}

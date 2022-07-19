package rgdb

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type driver struct {
	logger *zap.Logger
	cfg    *Config
	pool   *pgxpool.Pool
}

func New(logger *zap.Logger, cfg *Config) *driver {
	return &driver{
		logger: logger,
		cfg:    cfg,
	}
}

func (d *driver) Connect(ctx context.Context) error {
	pool, err := pgxpool.Connect(ctx, d.cfg.GetConnectionString())

	if err != nil {
		return err
	}

	err = pool.Ping(ctx)

	if err != nil {
		return nil
	}

	d.pool = pool

	return nil
}

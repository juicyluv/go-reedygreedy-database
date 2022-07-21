package rgdb

import (
	"context"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Client struct {
	logger *zap.Logger
	cfg    *Config
	pool   *pgxpool.Pool
}

func New(logger *zap.Logger, cfg *Config) (*Client, error) {
	config, err := pgxpool.ParseConfig(cfg.GetConnectionString())

	if err != nil {
		return nil, err
	}

	config.ConnConfig.Logger = zapadapter.NewLogger(logger)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	return &Client{
		logger: logger,
		cfg:    cfg,
		pool:   pool,
	}, nil
}

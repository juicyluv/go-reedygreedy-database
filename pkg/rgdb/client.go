package rgdb

import (
	"context"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/juicyluv/rgutils/pkg/logger"
)

type Client struct {
	logger *logger.Logger
	cfg    *Config
	pool   *pgxpool.Pool
}

func New(logger *logger.Logger, cfg *Config) (*Client, error) {
	config, err := pgxpool.ParseConfig(cfg.GetConnectionString())

	if err != nil {
		return nil, err
	}

	config.ConnConfig.Logger = zerologadapter.NewLogger(*logger.Logger)

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

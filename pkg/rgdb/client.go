package rgdb

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/juicyluv/rgutils/pkg/logger"
)

type Driver interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Close()
}

type Client struct {
	Logger *logger.Logger
	Config *Config
	Driver Driver
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
		Logger: logger,
		Config: cfg,
		Driver: pool,
	}, nil
}

func (c *Client) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return c.Driver.Query(ctx, sql, args)
}

func (c *Client) Close() {
	c.Driver.Close()
}

package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"

	"wallet/pkg/logger"
)

type (
	Client struct {
		*redis.Client
	}
)

func NewClient(url string) (client *Client, err error) {
	options, err := redis.ParseURL(url)

	if err != nil {
		err = errors.New("failed to parse URL for Redis: " + err.Error())
		return
	}

	redisClient := redis.NewClient(options)
	redis.SetLogger(dummyLogger{})

	if err := redisotel.InstrumentTracing(redisClient); err != nil {
		return nil, fmt.Errorf("failed to enable tracing for Redis: %v", err)
	}

	if err := redisotel.InstrumentMetrics(redisClient); err != nil {
		return nil, fmt.Errorf("failed to enable metrics for Redis: %v", err)
	}

	client = &Client{redisClient}
	ctx := context.Background()

	if client.Client.Ping(ctx) != nil {
		return nil, errors.New("failed to connect to the Redis: " + url)
	}
	logger.Log().Infof(ctx, "Redis successfully connect to %v", url)

	return client, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}

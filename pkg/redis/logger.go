package redis

import (
	"context"

	"wallet/pkg/logger"
)

type dummyLogger struct{}

func (d dummyLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	logger.Log().Infof(ctx, "from redis: "+format, v...)
}

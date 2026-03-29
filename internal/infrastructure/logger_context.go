package infrastructure

import (
	"context"

	"go.uber.org/zap"
)

type loggerKeyType struct{}

var loggerKey = loggerKeyType{}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetLogger(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return l
	}
	return zap.NewNop()
}

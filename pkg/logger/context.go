package logger

import (
	"context"
	"sync"

	"log/slog"
)

type ctxKey string

const (
	loggerCtxKey ctxKey = "LOGGER_CTX_KEY"
)

// loggerCtx private storage for logger's context.
type loggerCtx struct {
	mu    sync.RWMutex
	items map[string]any
}

// InitLoggerCtx init logger's context with storage.
func InitLoggerCtx(ctx context.Context) context.Context {
	logCtx := &loggerCtx{
		items: make(map[string]any),
	}
	return context.WithValue(ctx, loggerCtxKey, logCtx)
}

// SetItem set key+value in passed context for logger.
func SetItem(ctx context.Context, key string, value any) {
	logCtx, ok := ctx.Value(loggerCtxKey).(*loggerCtx)
	if !ok {
		return
	}

	logCtx.setItem(key, value)
}

// getAttrs return attrs log logger's context.
func getAttrs(ctx context.Context) []slog.Attr {
	logCtx, ok := ctx.Value(loggerCtxKey).(*loggerCtx)
	if !ok {
		return nil
	}

	return logCtx.getAttrs()
}

func (logCtx *loggerCtx) setItem(key string, value any) {
	logCtx.mu.Lock()
	defer logCtx.mu.Unlock()

	logCtx.items[key] = value
}

func (logCtx *loggerCtx) getAttrs() []slog.Attr {
	logCtx.mu.RLock()
	defer logCtx.mu.RUnlock()

	attrs := make([]slog.Attr, 0, len(logCtx.items))
	for k, v := range logCtx.items {
		attrs = append(attrs, slog.Any(k, v))
	}

	return attrs
}

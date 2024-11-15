package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

// init globalLogger with default slog.log.
var globalLogger = &Logger{l: slog.Default()} //nolint: gochecknoglobals // global by design

type Logger struct {
	l *slog.Logger
}

// Create create and return new logger.
func Create(config *Config) *Logger {
	var levelMapping = map[string]slog.Level{
		"debug": slog.LevelDebug,
		"info":  slog.LevelInfo,
		"warn":  slog.LevelWarn,
		"error": slog.LevelError,
	}

	level, ok := levelMapping[config.Level]
	if !ok {
		panic(fmt.Errorf("can't init logger with log level from config: %s", config.Level))
	}

	logLevel := &slog.LevelVar{}
	logLevel.Set(level)

	logger := &Logger{
		l: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})),
	}

	return logger
}

// InitGlobal init global logger for use as global.
func MustInitGlobal(config *Config) {
	globalLogger = Create(config)
}

// wo context

func (l *Logger) Debug(msg string, args ...any) {
	l.log(context.Background(), slog.LevelDebug, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log(context.Background(), slog.LevelInfo, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log(context.Background(), slog.LevelWarn, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log(context.Background(), slog.LevelError, msg, args...)
}

// with context

func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelDebug, msg, args...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelInfo, msg, args...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelWarn, msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log(ctx, slog.LevelError, msg, args...)
}

// common log method to log:
// * check if current log level is enabled - to prevent useless workflows around args.
func (l *Logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.l.Enabled(ctx, level) {
		return
	}

	l.l.Log(ctx, level, msg, getArgs(ctx, args...)...)
}

// global usage logger

// wo context

func Debug(msg string, args ...any) {
	globalLogger.DebugContext(context.Background(), msg, args...)
}

func Info(msg string, args ...any) {
	globalLogger.InfoContext(context.Background(), msg, args...)
}

func Warn(msg string, args ...any) {
	globalLogger.WarnContext(context.Background(), msg, args...)
}

func Error(msg string, args ...any) {
	globalLogger.ErrorContext(context.Background(), msg, args...)
}

// with context

func DebugContext(ctx context.Context, msg string, args ...any) {
	globalLogger.DebugContext(ctx, msg, args...)
}

func InfoContext(ctx context.Context, msg string, args ...any) {
	globalLogger.InfoContext(ctx, msg, args...)
}

func WarnContext(ctx context.Context, msg string, args ...any) {
	globalLogger.WarnContext(ctx, msg, args...)
}

func ErrorContext(ctx context.Context, msg string, args ...any) {
	globalLogger.ErrorContext(ctx, msg, args...)
}

// getArgs collects all args to log (include attrs and passed by client).
func getArgs(ctx context.Context, args ...any) []any {
	attrs := getAttrs(ctx)
	a := make([]any, 0, len(attrs)*2+len(args)) // attrs contains 2 value (key+value)
	for _, attr := range attrs {
		a = append(a, attr.Key, attr.Value.String())
	}
	if len(args) > 0 {
		a = append(a, args...)
	}

	return a
}

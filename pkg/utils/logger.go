package utils

import (
	"context"
	"fmt"
)

type contextKey string

var loggerKey = contextKey("logger")

type Logger interface {
	Printf(format string, v ...interface{})
}

func (c contextKey) String() string {
	return string(c)
}
func ContextWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerFromContext(ctx context.Context) (Logger, bool) {
	logger, ok := ctx.Value(loggerKey).(Logger)
	return logger, ok
}

type TestLogger struct {
	logs []string
}

func (t *TestLogger) Printf(format string, v ...interface{}) {
	t.logs = append(t.logs, fmt.Sprintf(format, v...))
}

func (t *TestLogger) Logs() []string {
	return t.logs
}

func NewTestLogger() *TestLogger {
	return &TestLogger{
		logs: []string{},
	}
}

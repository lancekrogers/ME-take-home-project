package utils

import (
	"context"
	"fmt"
	"log"
)

type contextKey string

const loggerKey contextKey = "logger"

func ContextWithLogger(ctx context.Context, logger *log.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func LoggerFromContext(ctx context.Context) (*log.Logger, bool) {
	logger, ok := ctx.Value(loggerKey).(*log.Logger)
	return logger, ok
}

type TestLogger struct {
	log.Logger
	messages []string
}

func NewTestLogger() *TestLogger {
	return &TestLogger{
		messages: []string{},
	}
}

func (tl *TestLogger) Printf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	tl.messages = append(tl.messages, message)
}

func (tl *TestLogger) Logs() []string {
	return tl.messages
}

func ContextWithTestLogger(ctx context.Context, logger *TestLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}


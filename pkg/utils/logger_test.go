package utils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithLogger(t *testing.T) {
	ctx := context.Background()
	logger := NewTestLogger()
	ctxWithLogger := ContextWithLogger(ctx, logger)
	assert.NotNil(t, ctxWithLogger.Value(loggerKey))
}

func TestLoggerFromContext(t *testing.T) {
	ctx := context.Background()
	logger := NewTestLogger()
	ctxWithLogger := ContextWithLogger(ctx, logger)

	retrievedLogger, ok := LoggerFromContext(ctxWithLogger)
	assert.True(t, ok)
	assert.Equal(t, logger, retrievedLogger)

	_, ok = LoggerFromContext(ctx)
	assert.False(t, ok)
}

func TestTestLogger(t *testing.T) {
	logger := NewTestLogger()
	logger.Printf("Test log %d", 1)
	logger.Printf("Another log %s", "message")
	logs := logger.Logs()
	assert.Len(t, logs, 2)
	assert.Equal(t, "Test log 1", logs[0])
	assert.Equal(t, "Another log message", logs[1])
}

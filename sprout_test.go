package sprout

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFunctionHandler_DefaultValues(t *testing.T) {
	handler := NewFunctionHandler()

	assert.NotNil(t, handler)
	assert.NotNil(t, handler.Logger)
}

func TestNewFunctionHandler_CustomValues(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	handler := NewFunctionHandler(
		WithLogger(logger),
	)

	assert.NotNil(t, handler)
	assert.Equal(t, logger, handler.Logger())
}

func TestWithLogger(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	option := WithLogger(logger)

	handler := NewFunctionHandler()
	option(handler) // Apply the option

	assert.Equal(t, logger, handler.Logger())
}

func TestWithParser(t *testing.T) {
	fnHandler := &DefaultHandler{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	option := WithHandler(fnHandler)

	handler := New()
	option(handler) // Apply the option

	assert.Equal(t, fnHandler, handler)
}

func TestWithNilHandler(t *testing.T) {
	fnHandler := &DefaultHandler{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
	option := WithHandler(nil)

	beforeApply := fnHandler
	option(beforeApply)

	assert.Equal(t, beforeApply, fnHandler)
}

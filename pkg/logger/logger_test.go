package logger_test

import (
	"testing"

	"github.com/goocarry/jenkinslocal/pkg/logger"
)

func TestLogger(t *testing.T) {
	t.Run("test log", func(t *testing.T) {
		logger := logger.Logger{}
		logger.Log("test log")
	})

	t.Run("test log", func(t *testing.T) {
		logger := logger.Logger{}
		logger.Log("test log 2")
	})
}

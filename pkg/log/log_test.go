package log_test

import (
	"errors"
	"testing"

	"github.com/redpkg/airship/pkg/log"
)

func TestInit(t *testing.T) {
	log.Init(log.Config{
		Level:   "info",
		Console: true,
	})

	log.Debug().Msg("debug message")
	log.Info().Msg("info message")
	log.Warn().Msg("warn message")
	log.Error().Msg("error message")
	log.Err(errors.New("foobar")).Msg("err message")
	// log.Fatal().Msg("fatal")
	// log.Panic().Msg("panic")
}

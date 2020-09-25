package log_test

import (
	"errors"
	"testing"

	"github.com/redpkg/airship/pkg/log"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	err := log.Init(log.Config{
		Level:   "debug",
		Console: true,
	})
	if !assert.NoError(err) {
		return
	}

	log.Debug().Msg("Debug")
	log.Info().Msg("Info")
	log.Warn().Msg("Warn")
	log.Error().Msg("Error")
	log.Err(errors.New("FOO")).Msg("Err")
	// log.Fatal().Msg("Fatal")
	// log.Panic().Msg("Panic")
}

package config_test

import (
	"fmt"
	"testing"

	"github.com/redpkg/airship/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	err := config.Init("../../config.yaml")
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", config.Box)
}

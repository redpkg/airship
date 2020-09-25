package redis_test

import (
	"testing"

	"github.com/redpkg/airship/pkg/redis"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	_, err := redis.New(redis.Config{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
	})
	assert.NoError(err)
}

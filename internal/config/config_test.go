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

	// fmt.Printf("%+v\n", config.Log)
	// fmt.Printf("%+v\n", config.Redis)
	// fmt.Printf("%+v\n", config.DB)
	// fmt.Println(config.ProjectPrefix)
	// fmt.Println(config.HTTPServerAddr)
	// fmt.Println(config.GRPCServerAddr)
	// fmt.Println(config.SessionTTL)
	// fmt.Println(config.TicketTTL)
	// fmt.Println(string(config.JWTSecret))
}

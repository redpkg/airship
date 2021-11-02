package main

import (
	"github.com/redpkg/airship/internal/config"
	"github.com/redpkg/airship/pkg/db"
	"github.com/redpkg/airship/pkg/log"
	"github.com/redpkg/airship/pkg/redis"
)

func main() {
	if err := config.Init("./config.yaml"); err != nil {
		panic(err)
	}
	if err := log.Init(config.Box.Log); err != nil {
		panic(err)
	}
	_, err := redis.New(config.Box.Redis)
	if err != nil {
		panic(err)
	}
	_, err = db.New(config.Box.DB)
	if err != nil {
		panic(err)
	}
}

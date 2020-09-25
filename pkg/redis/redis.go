package redis

import (
	"strconv"

	goredis "github.com/go-redis/redis/v8"
)

// Config redis config
type Config struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// New create redis instance
func New(conf Config) (*goredis.Client, error) {
	return goredis.NewClient(&goredis.Options{
		Addr:     conf.Host + ":" + strconv.Itoa(conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	}), nil
}

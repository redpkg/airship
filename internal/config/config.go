package config

import (
	"strings"
	"time"

	"github.com/redpkg/airship/pkg/db"
	"github.com/redpkg/airship/pkg/log"
	"github.com/redpkg/airship/pkg/redis"
	"github.com/spf13/viper"
)

// Box contained config for public access
var Box Config

// Config app config
type Config struct {
	Log            log.Config    `mapstructure:"log"`
	Redis          redis.Config  `mapstructure:"redis"`
	DB             db.Config     `mapstructure:"db"`
	ProjectPrefix  string        `mapstructure:"project_prefix"`
	HTTPServerAddr string        `mapstructure:"http_server_addr"`
	GRPCServerAddr string        `mapstructure:"grpc_server_addr"`
	JWTSecret      string        `mapstructure:"jwt_secret"`
	SessionTTL     time.Duration `mapstructure:"session_ttl"`
	TicketTTL      time.Duration `mapstructure:"ticket_ttl"`
}

// Init init config instance
func Init(files ...string) error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, file := range files {
		viper.SetConfigFile(file)
		if err := viper.MergeInConfig(); err != nil {
			return err
		}
	}

	if err := viper.Unmarshal(&Box); err != nil {
		return err
	}

	return nil
}

package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// Config database config
type Config struct {
	Writer          ConfigNode    `mapstructure:"writer"`
	Readers         []ConfigNode  `mapstructure:"readers"`
	Database        string        `mapstructure:"database"`
	Timezone        string        `mapstructure:"timezone"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
}

// ConfigNode database node config
type ConfigNode struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// New create db instance
func New(conf Config) (*gorm.DB, error) {
	db, err := gorm.Open(newDialector(conf.Writer, conf), &gorm.Config{
		DisableAutomaticPing: true,
	})
	if err != nil {
		return nil, err
	}

	readers := []gorm.Dialector{}
	for _, confReader := range conf.Readers {
		readers = append(readers, newDialector(confReader, conf))
	}

	if err := db.Use(
		dbresolver.Register(dbresolver.Config{
			Replicas: readers,
		}).
			SetConnMaxIdleTime(conf.ConnMaxIdleTime).
			SetConnMaxLifetime(conf.ConnMaxLifetime).
			SetMaxIdleConns(conf.MaxIdleConns).
			SetMaxOpenConns(conf.MaxOpenConns),
	); err != nil {
		return nil, err
	}

	return db, nil
}

func newDialector(confNode ConfigNode, conf Config) gorm.Dialector {
	return postgres.Open(buildDSN(confNode.Host, confNode.Port, confNode.Username, confNode.Password, conf.Database, conf.Timezone))
}

func buildDSN(host string, port int, username, password, database, timezone string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", host, port, username, password, database, timezone)
}

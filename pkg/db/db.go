package db

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

// Config database config
type Config struct {
	Writer          ConfigNode    `mapstructure:"writer"`
	Readers         []ConfigNode  `mapstructure:"readers"`
	Database        string        `mapstructure:"database"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	Timezone        string        `mapstructure:"timezone"`
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

	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: readers,
	}).
		SetConnMaxIdleTime(conf.ConnMaxIdleTime).
		SetConnMaxLifetime(conf.ConnMaxLifetime).
		SetMaxIdleConns(conf.MaxIdleConns).
		SetMaxOpenConns(conf.MaxOpenConns)); err != nil {
		return nil, err
	}

	return db, nil
}

func newDialector(confNode ConfigNode, conf Config) gorm.Dialector {
	return mysql.Open(buildDSN(confNode.Host, confNode.Port, confNode.Username, confNode.Password, conf.Database, conf.Timezone))
}

func buildDSN(host string, port int, username, password, database, timezone string) string {
	var s strings.Builder

	s.WriteString(username)
	s.WriteString(":")
	s.WriteString(password)
	s.WriteString("@tcp(")
	s.WriteString(host)
	s.WriteString(":")
	s.WriteString(strconv.Itoa(port))
	s.WriteString(")/")
	s.WriteString(database)
	s.WriteString("?parseTime=true&loc=")
	s.WriteString(url.QueryEscape(timezone))

	return s.String()
}

package db_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/redpkg/airship/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	db, err := db.New(db.Config{
		Writer: db.ConfigNode{
			Host:     "ybtv-dev.cluster-cewl1btfswol.ap-northeast-1.rds.amazonaws.com",
			Port:     3306,
			Username: "admin",
			Password: "95XSTYNFpvAYd2SJKh6PQwjwVaHa2T7F",
		},
		Readers: []db.ConfigNode{
			{
				Host:     "ybtv-dev.cluster-ro-cewl1btfswol.ap-northeast-1.rds.amazonaws.com",
				Port:     3306,
				Username: "admin",
				Password: "95XSTYNFpvAYd2SJKh6PQwjwVaHa2T7F",
			},
		},
		Database:        "ybtv",
		Timezone:        "UTC",
		ConnMaxIdleTime: time.Minute * 5,
		ConnMaxLifetime: time.Minute * 60,
		MaxIdleConns:    5,
		MaxOpenConns:    10,
	})
	if !assert.NoError(err) {
		return
	}

	result := map[string]interface{}{}
	db.Table("channels").Take(&result)
	fmt.Printf("%+v\n", result)
}

func Benchmark(b *testing.B) {
	db, err := db.New(db.Config{
		Writer: db.ConfigNode{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "root",
		},
		Readers: []db.ConfigNode{
			{
				Host:     "localhost",
				Port:     3306,
				Username: "root",
				Password: "root",
			},
		},
		Database:        "airship",
		Timezone:        "UTC",
		ConnMaxIdleTime: time.Minute * 5,
		ConnMaxLifetime: time.Minute * 60,
		MaxIdleConns:    5,
		MaxOpenConns:    10,
	})
	if err != nil {
		panic(err)
	}

	result := map[string]interface{}{}
	for i := 0; i < b.N; i++ {
		db.Table("users").Take(&result)
	}
}

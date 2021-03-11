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
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "secret",
		},
		Readers: []db.ConfigNode{
			{
				Host:     "localhost",
				Port:     5432,
				Username: "postgres",
				Password: "secret",
			},
		},
		Database:        "airship",
		Timezone:        "UTC",
		ConnMaxIdleTime: time.Minute * 5,
		ConnMaxLifetime: time.Hour,
		MaxIdleConns:    5,
		MaxOpenConns:    10,
	})
	if !assert.NoError(err) {
		return
	}

	fmt.Printf("%+v\n", db)
}

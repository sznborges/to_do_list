package ioc

import (
	"database/sql"
	"sync"

	"github.com/sznborges/to_do_list/config"
	"github.com/sznborges/to_do_list/db"
)

var (
	dbOnce       sync.Once
	dbi          *sql.DB
)


func DB() *sql.DB {
	dbOnce.Do(func() {
		d, err := db.OpenConnection(db.Config{
			Host:              config.GetString("DB_HOST"),
			Port:              config.GetString("DB_PORT"),
			User:              config.GetString("DB_USER"),
			Password:          config.GetString("DB_PASSWORD"),
			Name:              config.GetString("DB_NAME"),
			PoolSize:          config.GetInt("DB_POOL_SIZE"),
			ConnMaxTTL:        config.GetDuration("DB_CONN_MAX_TTL_MILLIS"),
			TimeoutSeconds:    config.GetInt("DB_TIMEOUT_SECONDS"),
			LockTimeoutMillis: config.GetInt("DB_LOCK_TIMEOUT_MILLIS"),
		})
		if err != nil {
			panic(err)
		}
		dbi = d
	})
	return dbi
}
package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/sznborges/to_do_list/config"
	"github.com/sznborges/to_do_list/infra/logger"
)

var (
	postgresOnce   sync.Once
	clientInstance *sql.DB
)

type PostgresConnection struct {
	connectionString string
}

func NewPostgresConnection() *PostgresConnection {
	str := &PostgresConnection{}
	str.connectionString = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.GetString("DB_HOST"),
		config.GetString("DB_USER"),
		config.GetString("DB_PASSWORD"),
		config.GetString("DB_NAME"),
		config.GetInt("DB_PORT"),
		config.GetString("DB_SSLMODE"),
	)
	return str
}

func (c *PostgresConnection) GetConnection() *sql.DB {
	postgresOnce.Do(func() {
		conn, err := sql.Open("postgres", c.connectionString)
		if err != nil {
			logger.Logger.Fatal(err)
			return
		}
		clientInstance = conn
	})
	return clientInstance
}
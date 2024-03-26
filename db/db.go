package db

import (
	"database/sql"
	"fmt"
	"time"

	pgx "github.com/jackc/pgx/v4/stdlib" // postgresql driver
	"github.com/sznborges/to_do_list/config"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)
	
type Config struct {
		Host              string
		Port              string
		User              string
		Password          string
		Name              string
		PoolSize          int
		ConnMaxTTL        time.Duration
		TimeoutSeconds    int
		LockTimeoutMillis int
	}
	
	func (c Config) dsn() string {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=%d statement_timeout=%ds lock_timeout=%d",
			c.Host, c.Port, c.User, c.Password, c.Name, c.TimeoutSeconds, c.TimeoutSeconds, c.LockTimeoutMillis)
	}
	
	func OpenConnection(c Config) (*sql.DB, error) {
		service := config.GetString("SERVICE_NAME")
		sqltrace.Register("pgx", &pgx.Driver{}, sqltrace.WithServiceName(fmt.Sprint(service, "-db")))
		db, err := sqltrace.Open("pgx", c.dsn())
		if err != nil {
			return nil, err
		}
		db.SetMaxIdleConns(c.PoolSize)
		db.SetMaxOpenConns(c.PoolSize)
		db.SetConnMaxIdleTime(c.ConnMaxTTL)
		err = db.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil
	}


package domain

import "database/sql"

type DatabaseConnector interface {
	GetConnection() *sql.DB
}
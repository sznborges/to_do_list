package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
    port     = 5432
    user     = "to_do_user"
    password = "to_do_password"
    dbname   = "to_do"
)

func connectDB() *sql.DB {
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Erro ao conectar ao banco de dados:", err)
    }
    return db
}


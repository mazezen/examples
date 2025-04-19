package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

type Database struct {
	Conn   *sql.DB
	Logger zerolog.Logger
}

func Init(logger zerolog.Logger) (Database, error) {
	db := Database{}
	dsn := "postgres://letterpress:letterpress_secrets@postgres:5432/letterpress_db?sslmode=disable"
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	db.Logger = logger
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}

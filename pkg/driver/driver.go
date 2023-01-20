package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	maxOpenDBConn = 10
	maxIdleDBConn = 5
	maxDBLifetime = 5 * time.Minute
)

type DB struct {
	SQL *sql.DB
}

// ConnectDB creates connection pool for Postgres
func ConnectDB(dsn string) *DB {
	dbPool, err := NewDB(dsn)
	if err != nil {
		panic(err)
	}
	dbPool.SetMaxOpenConns(maxOpenDBConn)
	dbPool.SetConnMaxIdleTime(maxIdleDBConn)
	dbPool.SetConnMaxLifetime(maxDBLifetime)
	dbConn := &DB{
		SQL: dbPool,
	}
	return dbConn
}

// NewDB creates a new database for app
func NewDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

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

var dbConn = &DB{}

// ConnectDB creates connection pool for Postgres
func ConnectDB(dsn string) (*DB, error) {
	dbPool, err := NewDB(dsn)
	if err != nil {
		panic(err)
	}
	dbPool.SetMaxOpenConns(maxOpenDBConn)
	dbPool.SetConnMaxIdleTime(maxIdleDBConn)
	dbPool.SetConnMaxLifetime(maxDBLifetime)

	if err = testDB(dbPool); err != nil {
		return nil, err
	}
	dbConn.SQL = dbPool
	return dbConn, nil
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

// testDB tests the database connection
func testDB(d *sql.DB) error {
	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}

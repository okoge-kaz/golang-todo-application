package db

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // initialize mysql driver
	"github.com/jmoiron/sqlx"
)

var _db *sqlx.DB

// DefaultDSN creates default DSN string
func defaultDSN(host, port, user, password, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo", user, password, host, port, dbName)
}

// connect to database
func connect(dsn string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Init initializes the database connection
func Init(host, port, user, password, dbName string) error {
	// initialize database connection
	dsn := defaultDSN(host, port, user, password, dbName)

	// connect to database
	conn, err := connect(dsn)
	if err != nil {
		log.Fatal(err)
	}

	// check connection
	err = conn.Ping()
	if err != nil {
		return err
	}

	// bind database connection
	_db = conn

	return nil
}

// Disconnect closes connection
func Disconnect() {
	if _db != nil {
		_db.Close()
	}
}

// GetConnection returns DB connection
func GetConnection() (*sqlx.DB, error) {
	if _db != nil {
		return _db, nil
	}
	return nil, errors.New("connection is not established")
}

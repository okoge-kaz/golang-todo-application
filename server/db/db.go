package db

import (
	"errors"
	"fmt"

	"github.com/okoge-kaz/golang-todo-application/server/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

// DefaultDSN creates default DSN string
func defaultDSN(host, port, user, password, dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo", user, password, host, port, dbName)
}

// connect to database
func connect(dsn string) (*gorm.DB, error) {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
		panic("failed to connect database")
	}

	// bind database connection
	_db = conn

	// create tables (if not exists)
	createTables(_db)

	// auto migration
	autoMigrate(_db)

	return nil
}

// GetConnection returns DB connection
func GetConnection() (*gorm.DB, error) {
	if _db != nil {
		return _db, nil
	}
	return nil, errors.New("connection is not established")
}

// for migration
func autoMigrate(_db *gorm.DB) {
	// migration tables
	_db.AutoMigrate(&entities.Task{})
	_db.AutoMigrate(&entities.User{})
	_db.AutoMigrate(&entities.Ownership{})
	_db.AutoMigrate(&entities.Category{})
}

// create tables
func createTables(_db *gorm.DB) {
	if !_db.Migrator().HasTable(&entities.Task{}) {
		_db.Migrator().CreateTable(&entities.Task{})
	}
	if !_db.Migrator().HasTable(&entities.User{}) {
		_db.Migrator().CreateTable(&entities.User{})
	}
	if !_db.Migrator().HasTable(&entities.Ownership{}) {
		_db.Migrator().CreateTable(&entities.Ownership{})
	}
	if !_db.Migrator().HasTable(&entities.Category{}) {
		_db.Migrator().CreateTable(&entities.Category{})
	}
}

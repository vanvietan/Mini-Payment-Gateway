package data

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var dbConn *gorm.DB

const (
	user     = "postgres"
	password = "admin"
	port     = 5432
	dbname   = "postgres"
	host     = "localhost"
)

func CreateDatabaseConnection() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user =%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("Database Connected!")

	// create the connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbConn = db

	return nil
}
func GetDatabaseConnection() (*gorm.DB, error) {
	if err := CreateDatabaseConnection(); err != nil {
		return nil, err
	}
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn, err
	}
	return dbConn, nil
}

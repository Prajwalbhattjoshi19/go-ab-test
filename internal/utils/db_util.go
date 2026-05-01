package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	ora "github.com/sijms/go-ora/v2"
)

// OpenDB opens a database connection using the provided driver and data source name (DSN).
func OpenDB(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}

func GetDBDriver() *sql.DB {
	dbPort, err := strconv.Atoi(os.Getenv("db_port"))
	if err != nil {
		log.Fatal("Invalid db_port value: ", err)
	}
	connStr := ora.BuildUrl(os.Getenv("db_server"), dbPort, os.Getenv("db_service_name"), os.Getenv("db_username"), os.Getenv("db_password"), nil)
	conn, err := sql.Open("oracle", connStr)
	conn.SetMaxIdleConns(50)
	conn.SetMaxOpenConns(50)
	conn.SetConnMaxLifetime(0)
	conn.SetConnMaxIdleTime(0)
	// check for error
	if err != nil {
		log.Fatal("Failed to create database connection: ", err)
		panic(err)
	}
	// check for error
	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
		panic(err)
	}
	//ping the database to check if the connection is alive
	err = conn.Ping()
	if err != nil {
		log.Fatal("Failed to ping database: ", err)
		panic(err)
	}
	return conn
}

// 代码生成时间: 2025-10-11 02:40:23
package main

import (
    "database/sql"
    "fmt"
    "log"
    "strings"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

// DatabaseConfig contains the configuration for database connection
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// DatabasePool defines the structure to manage database connection pool
type DatabasePool struct {
    *sql.DB
    Config DatabaseConfig
}

// NewDatabasePool creates a new database connection pool
func NewDatabasePool(config DatabaseConfig) (*DatabasePool, error) {
    // Construct the DSN (Data Source Name)
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        config.User,
        config.Password,
        config.Host,
        config.Port,
        config.DBName,
    )

    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Set the maximum number of connections in the idle connection pool
    db.SetMaxIdleConns(10)

    // Set the maximum number of open connections to the database
    db.SetMaxOpenConns(100)

    // Set the connection max lifetime
    db.SetConnMaxLifetime(3600)

    // Ping the database to verify the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return &DatabasePool{DB: db, Config: config}, nil
}

// Close the database pool connection
func (dp *DatabasePool) Close() error {
    return dp.DB.Close()
}

func main() {
    // Define the database configuration
    config := DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "root",
        Password: "password",
        DBName:   "testdb",
    }

    // Create a new database pool
    dbPool, err := NewDatabasePool(config)
    if err != nil {
        log.Fatalf("Failed to create database pool: %v", err)
    }
    defer dbPool.Close()

    // Example of using the database pool to execute a query
    var version string
    err = dbPool.QueryRow("SELECT VERSION()").Scan(&version)
    if err != nil {
        log.Printf("Failed to execute query: %v", err)
    } else {
        fmt.Println("Database version: ", version)
    }
}

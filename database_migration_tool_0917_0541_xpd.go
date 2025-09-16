// 代码生成时间: 2025-09-17 05:41:00
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/logger"
    "github.com/kataras/iris/v12/middleware/recover"
    "github.com/go-pg/migrations/v8"
    "github.com/go-pg/pg/v10"
)

// databaseConfig holds information about the database
type databaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Database string
}

// MigrationTool represents the database migration tool
type MigrationTool struct {
    db     *pg.DB
    options migrations.MigrationOptions
}

func main() {
    // Initialize Iris
    app := iris.Default()
    app.Use(recover.New())
    app.Use(logger.New())

    // Initialize database configuration
    dbConfig := databaseConfig{
        Host:     "localhost",
        Port:     5432,
        User:     "postgres",
        Password: "password",
        Database: "postgres",
    }

    // Create a new database migration tool
    migrationTool := newMigrationTool(dbConfig)

    // Define routes for the migration tool
    app.Get("/migrate", migrationTool.migrate)
    app.Get("/rollback", migrationTool.rollback)

    // Start the Iris server
    log.Printf("Server is running on :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

// newMigrationTool initializes and returns a new MigrationTool
func newMigrationTool(dbConfig databaseConfig) *MigrationTool {
    // Create a new database connection
    db := pg.Connect(&pg.Options{
        User:     dbConfig.User,
        Password: dbConfig.Password,
        Database: dbConfig.Database,
        Addr:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
    })

    // Create a new migration tool with the database connection
    return &MigrationTool{
        db: db,
        options: migrations.MigrationOptions{
            TableName: "migrations",
        },
    }
}

// migrate applies all pending database migrations
func (m *MigrationTool) migrate(ctx iris.Context) {
    if err := migrations.Run(m.db, m.options); err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("There was an error during migration: " + err.Error())
        log.Printf("Error during migration: %v", err)
        return
    }

    ctx.WriteString("Migration completed successfully")
}

// rollback rolls back the latest database migration
func (m *MigrationTool) rollback(ctx iris.Context) {
    if err := migrations.Rollback(m.db, m.options); err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("There was an error during rollback: " + err.Error())
        log.Printf("Error during rollback: %v", err)
        return
    }

    ctx.WriteString("Rollback completed successfully")
}

// 代码生成时间: 2025-08-21 00:07:39
package main

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/sirupsen/logrus"
)

// Logger is a wrapper for the logrus.Logger to handle log operations.
type Logger struct {
    *logrus.Logger
}

// NewLogger creates a new Logger instance.
func NewLogger() *Logger {
    log := logrus.New()
    log.Formatter = &logrus.TextFormatter{
        ForceColors:     true,
        FullTimestamp:   true,
        TimestampFormat: time.RFC3339Nano,
    }
    return &Logger{Logger: log}
}

// AuditLogHandler is a middleware that logs audit events.
func AuditLogHandler(ctx iris.Context) {
    // Start timing the request.
    start := time.Now()

    // Proceed to the next handler in the chain.
    ctx.Next()

    // Calculate the duration of the request.
    duration := time.Since(start)

    // Log the audit event.
    logger := NewLogger()
    logger.Infof("Audit Event:
	Method: %s
	Path: %s
	Duration: %v
	IP: %s", ctx.Method(), ctx.Path(), duration, ctx.RemoteAddr())
}

// FileAuditLogger is a function that writes audit logs to a file.
func FileAuditLogger(path string) (*os.File, error) {
    // Ensure the directory exists.
    if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
        return nil, err
    }

    // Open the file for appending.
    file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }
    return file, nil
}

func main() {
    app := iris.New()

    // Set up a file for audit logs.
    logFile, err := FileAuditLogger("./audit.log")
    if err != nil {
        fmt.Printf("Failed to create log file: %v
", err)
        return
    }
    defer logFile.Close()

    // Set up the audit log middleware.
    app.Use(AuditLogHandler)

    // Example route that triggers an audit log.
    app.Get("/example", func(ctx iris.Context) {
        ctx.JSON(iris.StatusOK, iris.Map{
            "message": "Hello from Audit Log Example!",
        })
    })

    // Start the server.
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Printf("Failed to start server: %v
", err)
    }
}

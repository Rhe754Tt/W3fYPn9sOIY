// 代码生成时间: 2025-09-23 22:27:47
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// ApiResponse represents a structured API response
type ApiResponse struct {
    Timestamp time.Time `json:"timestamp"`
    Status    string    `json:"status"`
    Data      interface{} `json:"data"`
    Message   string    `json:"message,omitempty"`
    Error     *ErrorInfo `json:"error,omitempty"`
}

// ErrorInfo contains error details
type ErrorInfo struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// FormatError formats an error as an ApiResponse
func FormatError(code int, message string) *ApiResponse {
    return &ApiResponse{
        Timestamp: time.Now(),
        Status:    "error",
        Error:     &ErrorInfo{Code: code, Message: message},
    }
}

// FormatSuccess formats a success response as an ApiResponse
func FormatSuccess(data interface{}, message string) *ApiResponse {
    return &ApiResponse{
        Timestamp: time.Now(),
        Status:    "success",
        Data:      data,
        Message:   message,
    }
}

func main() {
    app := iris.New()

    // Define a route for a sample API
    app.Get("/api/sample", func(ctx iris.Context) {
        // Simulate a successful response
        response := FormatSuccess("Hello, IRIS!", "This is a sample success response.")
        ctx.JSON(http.StatusOK, response)
    })

    // Define a route for an error scenario
    app.Get("/api/error", func(ctx iris.Context) {
        // Simulate an error response
        response := FormatError(http.StatusInternalServerError, "Internal Server Error")
        ctx.JSON(http.StatusInternalServerError, response)
    })

    // Start the server
    log.Printf("Server is running at :8080")
    if err := app.Listen(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

// 代码生成时间: 2025-09-05 11:49:00
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/kataras/iris/v12"
)

// Main function to run the application
func main() {
    app := iris.New()

    // Define a route for performance testing
    app.Get("/test", func(ctx iris.Context) {
        start := time.Now()
        // Simulate a long-running operation
        time.Sleep(100 * time.Millisecond)
        fmt.Println("Performing long-running operation...")
        duration := time.Since(start)
        ctx.WriteString("Operation completed in " + duration.String())
    })

    // Start the server with a custom error handler
    log.Fatal(app.Listen(":8080", iris.WithOptimizations()))
}

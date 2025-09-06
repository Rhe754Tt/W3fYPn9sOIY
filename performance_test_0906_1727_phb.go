// 代码生成时间: 2025-09-06 17:27:01
package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/kataras/iris/v12"
)

// Main is the entry point for the performance test application.
func main() {
    app := iris.New()

    // Define a route for a simple performance test.
    // This route will be used to measure the response time and throughput.
    app.Get("/performance", func(ctx iris.Context) {
        // Simulate some processing time.
        time.Sleep(10 * time.Millisecond)
        ctx.WriteString("Hello from Iris!")
    })

    // Define a route for starting a performance test.
    app.Get("/start-performance-test", func(ctx iris.Context) {
        startPerformanceTest(ctx)
    })

    // Start the Iris server.
    // The ListenAndServe method will block until the server stops.
    if err := app.Listen(":8080", iris.WithOptimizations()); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}

// startPerformanceTest simulates a performance test by sending multiple requests to the /performance route.
// It prints out the total time taken and the number of requests per second.
func startPerformanceTest(ctx iris.Context) {
    const numRequests = 100
    const duration = 5 * time.Second
    const requestURL = "http://localhost:8080/performance"

    start := time.Now()
    var totalBytes int64

    // Create a HTTP client for sending requests.
    client := &http.Client{}
    for i := 0; i < numRequests; i++ {
        req, err := http.NewRequest(http.MethodGet, requestURL, nil)
        if err != nil {
            ctx.WriteString("Error creating request: " + err.Error())
            return
        }

        // Send the request and measure the response time.
        resp, err := client.Do(req)
        if err != nil {
            ctx.WriteString("Error sending request: " + err.Error())
            return
        }
        defer resp.Body.Close()

        // Read the response body to ensure the full request cycle is measured.
        if _, err := io.Copy(io.Discard, resp.Body); err != nil {
            ctx.WriteString("Error reading response body: " + err.Error())
            return
        }

        totalBytes += resp.ContentLength
    }

    // Calculate the total time taken and the requests per second.
    elapsed := time.Since(start)
    requestsPerSecond := float64(numRequests) / elapsed.Seconds()

    // Write the results to the Iris context.
    ctx.JSON(iris.StatusOK, map[string]interface{}{
        "total_time": elapsed.String(),
        "requests_per_second": requestsPerSecond,
        "bytes_transferred": totalBytes,
    })
}

// 代码生成时间: 2025-10-02 03:26:22
// end_to_end_test.go

package main

import (
    "fmt"
    "log"
    "net/http"
    "testing"
    "github.com/kataras/iris/v12/httptest"
)

// setupApp sets up a simple Iris application for testing.
func setupApp() *httptest.Application {
    app := httptest.New(t, httptest.Debug(true))
    // Setup your Iris handlers here
    app.Get("/", func(ctx *iris.Context) {
        ctx.HTML(`<h1>Hello Iris</h1>`)
    })
    return app
}

// TestEndToEnd performs an end-to-end test for the Iris application.
func TestEndToEnd(t *testing.T) {
    app := setupApp()
    defer app.Close()

    // Test the root path
    e := app.Request("GET", "/").Expect().Status(http.StatusOK).Body()
    if !e.Equal(`<h1>Hello Iris</h1>`) {
        t.Fatalf("Expected body <h1>Hello Iris</h1>, but got %s", e)
    }
}

func main() {
    // The main function is just for running the app, not for testing.
    // Start your Iris application here if you want to run it standalone.
    // iris.Listen(":8080")
}

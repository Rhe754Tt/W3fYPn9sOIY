// 代码生成时间: 2025-08-24 04:08:44
package main

import (
    "encoding/json"
    "fmt"
    "github.com/kataras/iris/v12"
    "net/http"
)

// JSONConverter is a struct that holds the Iris application.
type JSONConverter struct {
    app *iris.Application
}

// NewJSONConverter creates and returns a new JSONConverter instance.
func NewJSONConverter() *JSONConverter {
    return &JSONConverter{
        app: iris.Default(),
    }
}

// Start starts the HTTP server with the defined routes.
func (j *JSONConverter) Start(port string) {
    j.app.Get("/convert", j.convertJSON)
    j.app.Listen(fmt.Sprintf(":%s", port))
}

// convertJSON handles the JSON conversion endpoint.
// It accepts a JSON payload, converts it, and returns the result.
func (j *JSONConverter) convertJSON(ctx iris.Context) {
    var payload map[string]interface{}

    // Decode the JSON payload into the payload map.
    if err := ctx.ReadJSON(&payload); err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": "Invalid JSON payload",
        })
        return
    }

    // Convert the map to a JSON string.
    jsonString, err := json.Marshal(payload)
    if err != nil {
        ctx.StatusCode(http.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": "Failed to convert JSON",
        })
        return
    }

    // Return the converted JSON string.
    ctx.JSON(iris.Map{
        "converted": string(jsonString),
    })
}

// main is the entry point for the program.
func main() {
    converter := NewJSONConverter()
    converter.Start("8080")
}

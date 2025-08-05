// 代码生成时间: 2025-08-05 18:27:05
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
    "github.com/nfnt/resize"
)

// Define the dimensions for the resized image
const (
    defaultWidth  = 800
    defaultHeight = 600
)

func main() {
    app := iris.New()
    app.Use(recover.New()) // Use recover middleware

    // Define the upload route
    app.Post("/upload", func(ctx iris.Context) {
        // Check if there is a file in the request
        if ctx.FormValue("image") == "" {
            ctx.Values().Set("error", "No image provided")
            ctx.JSON(http.StatusBadRequest, iris.Map{
                "error": "No image provided",
            })
            return
        }

        // Retrieve the uploaded file
        file, info, err := ctx.FormFile("image")
        if err != nil {
            ctx.Values().Set("error", err.Error())
            ctx.JSON(http.StatusInternalServerError, iris.Map{
                "error": err.Error(),
            })
            return
        }
        defer file.Close()

        // Save the file to a temporary location
        tempFile, err := ioutil.TempFile("
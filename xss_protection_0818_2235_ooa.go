// 代码生成时间: 2025-08-18 22:35:31
package main

import (
    "fmt"
    "net/http"
    "html"
    "github.com/kataras/iris/v12"
)

// sanitizeInput sanitizes the input to protect against XSS attacks.
func sanitizeInput(input string) string {
    // Use html.EscapeString to escape the HTML special characters.
    return html.EscapeString(input)
}

func main() {
    app := iris.New()

    // Define a route that handles a GET request to /xss.
    app.Get("/xss", func(ctx iris.Context) {
        // Retrieve the input value from the query parameter.
        input := ctx.URLParam("input")

        // Sanitize the input to prevent XSS attacks.
        sanitizedInput := sanitizeInput(input)

        // Render a template with the sanitized input.
        // This assumes you have a template file named `xss.html` in your 'views' directory.
        // The template should safely display the sanitized input.
        err := ctx.View("xss.html", iris.Map{
            "Input": sanitizedInput,
        })

        // Handle any errors that may occur during template rendering.
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            fmt.Fprintf(ctx, "An error occurred: %s", err.Error())
            return
        }
    })

    // Start the IRIS web server.
    // The server will listen on the default port (8080).
    app.Listen(":8080")
}

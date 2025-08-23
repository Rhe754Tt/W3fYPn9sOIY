// 代码生成时间: 2025-08-23 18:57:19
package main

import (
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// sanitizeInput removes potentially dangerous characters to prevent XSS attacks.
func sanitizeInput(input string) string {
    return strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", """, "&quot;", "'", "&#39;").Replace(input)
}

func main() {
    app := iris.New()
    
    // Define a route that echoes back user input, sanitized to prevent XSS.
    app.Get("/echo", func(ctx iris.Context) {
        input := ctx.URLParam("input")
        // Sanitize the input to prevent XSS attacks.
        sanitizedInput := sanitizeInput(input)
        ctx.WriteString("You said: " + sanitizedInput)
    })

    // Error handling for any unexpected issues.
    if err := app.Listen(":8080"); err != nil {
        panic(err)
    }
}

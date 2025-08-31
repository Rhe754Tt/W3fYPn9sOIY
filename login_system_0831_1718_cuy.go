// 代码生成时间: 2025-08-31 17:18:43
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/hero"
)

// User represents the user model for our application.
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginResponse represents the response after a successful login.
type LoginResponse struct {
    Message string `json:"message"`
}

func main() {
    app := iris.New()

    // Define a route that listens on POST method for login requests.
    app.Post("/login", func(ctx iris.Context) {
        // Decode the incoming request body into a User struct.
        var user User
        if err := ctx.ReadJSON(&user); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "Invalid request body."
            })
            return
        }

        // Simulate user verification logic.
        if user.Username != "admin" || user.Password != "password123" {
            ctx.StatusCode(iris.StatusUnauthorized)
            ctx.JSON(iris.Map{
                "error": "Invalid username or password."
            })
            return
        }

        // If user is valid, return a success response.
        ctx.StatusCode(iris.StatusOK)
        ctx.JSON(LoginResponse{Message: "Login successful."})
    })

    // Start the web server.
    fmt.Println("Server is running on :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Printf("An error occurred: %s
", err.Error())
    }
}

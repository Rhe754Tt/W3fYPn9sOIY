// 代码生成时间: 2025-08-10 22:05:19
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
)

// User represents the structure of a user for login purposes.
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func main() {
    // Initialize Iris
    app := iris.New()
    app.Use(recover.New()) // Recover middleware to handle panics

    // Define a route for login
    app.Post("/login", func(ctx iris.Context) {
        var user User
        if err := ctx.ReadJSON(&user); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": fmt.Sprintf("Error reading user data: %s", err),
            })
            return
        }

        // Here you would normally check the user's credentials against a database
        // For demonstration purposes, we're just checking if the username is 'admin' and the password is 'password123'
        if user.Username == "admin" && user.Password == "password123" {
            ctx.JSON(iris.Map{
                "message": "Login successful",
            })
        } else {
            ctx.StatusCode(iris.StatusUnauthorized)
            ctx.JSON(iris.Map{
                "error": "Invalid username or password",
            })
        }
    })

    // Start the server
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Printf("Server failed to start: %s", err)
    }
}

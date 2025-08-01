// 代码生成时间: 2025-08-02 00:59:52
package main

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// UserRole represents the type of user roles
type UserRole int

const (
    // UserRoleGuest represents a guest user
    UserRoleGuest UserRole = iota
    // UserRoleUser represents a regular user
    UserRoleUser
    // UserRoleAdmin represents an admin user
    UserRoleAdmin
)

// AccessControlMiddleware checks if the user is authorized to access the requested resource
func AccessControlMiddleware(allowedRole UserRole) iris.HandlerFunc {
    return func(ctx iris.Context) {
        // Get the user role from the request context
        userRole := UserRole(ctx.Values().GetInt("userRole"))
        // Check if the user has the required role
        if userRole < allowedRole {
            ctx.StatusCode(iris.StatusForbidden)
            ctx.Writef("Access denied. You don't have enough permissions.")
            return
        }
        // Continue to the next handler if the user is authorized
        ctx.Next()
    }
}

func main() {
    app := iris.New()

    // Define routes with access control
    app.Get("/admin", AccessControlMiddleware(UserRoleAdmin), func(ctx iris.Context) {
        ctx.Writef("Welcome, Admin!")
    })
    app.Get("/user", AccessControlMiddleware(UserRoleUser), func(ctx iris.Context) {
        ctx.Writef("Welcome, User!")
    })
    app.Get("/guest", AccessControlMiddleware(UserRoleGuest), func(ctx iris.Context) {
        ctx.Writef("Welcome, Guest!")
    })

    // Start the server
    app.Listen(":8080")
}

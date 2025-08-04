// 代码生成时间: 2025-08-05 07:05:29
Features:
- Code structure is clear and easy to understand.
- Includes appropriate error handling.
- Has necessary comments and documentation.
- Follows GOLANG best practices.
- Ensures code maintainability and scalability.
*/

package main

import (
    "fmt"
    "log"
    "strings"
    "github.com/kataras/iris/v12"
)

// AccessControl struct to store user roles
type AccessControl struct {
    allowedRoles []string
}

// NewAccessControl creates a new AccessControl instance
func NewAccessControl(allowedRoles []string) *AccessControl {
    return &AccessControl{allowedRoles: allowedRoles}
}

// Middleware to check if a user has the required role
func (ac *AccessControl) Middleware() iris.Handler {
    return func(ctx iris.Context) {
        userRole := ctx.Values().GetString("role")
        if userRole == "" {
            ctx.StatusCode(iris.StatusUnauthorized)
            ctx.WriteString("Unauthorized: User role is missing")
            return
        }
        for _, role := range ac.allowedRoles {
            if strings.EqualFold(role, userRole) {
                return
            }
        }
        ctx.StatusCode(iris.StatusForbidden)
        ctx.WriteString("Forbidden: User does not have the required role")
    }
}

func main() {
    app := iris.New()

    // Define allowed roles
    allowedRoles := []string{"admin", "moderator"}

    // Create an AccessControl instance
    ac := NewAccessControl(allowedRoles)

    // Register a middleware for access control
    app.Use(ac.Middleware())

    // Define a route that requires admin or moderator role
    app.Get("/protected", func(ctx iris.Context) {
        ctx.WriteString("Access granted to protected resource")
    })

    // Define a route that does not require any role
    app.Get("/public", func(ctx iris.Context) {
        ctx.WriteString("Access granted to public resource")
    })

    // Start the server
    log.Fatal(app.Listen(":8080"))
}

// 代码生成时间: 2025-08-26 23:43:31
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/middleware/recover"
    "github.com/kataras/iris/v12/middleware/logger"
    "github.com/kataras/iris/v12/middleware/cors"
)

// User represents the structure of a user
type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Role     string `json:"role"`
}

// Permission represents the structure of a permission
type Permission struct {
    ID         uint   `json:"id"`
    Permission string `json:"permission"`
    Role       string `json:"role"`
}

// UserPermissions is a struct that holds the mapping of users and their permissions
var UserPermissions map[uint]map[string]bool

// NewUser creates a new user
func NewUser(ctx iris.Context) {
    user := User{}
    if err := ctx.ReadJSON(&user); err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
        return
    }
    // Here you would typically save the user to a database
    fmt.Printf("New User: %+v\
", user)
    ctx.JSON(iris.Map{
        "message": "User created successfully",
        "user": user,
    })
}

// GetUserPermissions gets the permissions for a user
func GetUserPermissions(ctx iris.Context) {
    id := ctx.URLParam("id")
    if permissions, exists := UserPermissions[uint(id)]; exists {
        ctx.JSON(iris.Map{
            "id":        uint(id),
            "permissions": permissions,
        })
    } else {
        ctx.StatusCode(iris.StatusNotFound)
        ctx.JSON(iris.Map{
            "error": "User not found",
        })
    }
}

// AddPermissionToUser adds a permission to a user
func AddPermissionToUser(ctx iris.Context) {
    id := ctx.URLParam("id\)
    permission := ctx.URLParam("permission\)
    if _, exists := UserPermissions[uint(id)]; !exists {
        ctx.StatusCode(iris.StatusNotFound)
        ctx.JSON(iris.Map{
            "error": "User not found",
        })
        return
    }
    UserPermissions[uint(id)][permission] = true
    ctx.JSON(iris.Map{
        "message": "Permission added successfully",
    })
}

func main() {
    // Initialize Iris
    app := iris.New()
    app.Use(recover.New())
    app.Use(logger.New())
    app.Use(cors.New())

    // Define routes
    app.Post("/users", NewUser)
    app.Get("/users/{id}/permissions", GetUserPermissions)
    app.Put("/users/{id}/permissions/{permission}", AddPermissionToUser)

    // Start the server
    app.Listen(":8080")
}

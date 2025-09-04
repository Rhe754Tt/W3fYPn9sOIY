// 代码生成时间: 2025-09-04 20:04:03
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// App 结构体包含了Iris框架的引擎实例
type App struct {
    *iris.Application
}

// NewApp 创建一个新的App实例
func NewApp() *App {
    return &App{iris.New()}
}

// SetupRoutes 设置路由和处理函数
func (a *App) SetupRoutes() {
    // 定义一个GET请求的路由，用于返回一个简单的欢迎信息
    a.Get("/", func(ctx iris.Context) {
        ctx.WriteString("Welcome to the RESTful API!")
    })

    // 定义一个GET请求的路由，用于获取用户列表
    a.Get("/users", func(ctx iris.Context) {
        ctx.JSON(iris.StatusOK, []string{"User 1", "User 2"})
    })

    // 定义一个POST请求的路由，用于创建用户
    a.Post("/users", func(ctx iris.Context) {
        // 从请求体中解析用户数据
        var user struct{
            Name string `json:"name"`
        }
        if err := ctx.ReadJSON(&user); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.WriteString(fmt.Sprintf("Error reading request body: %s", err.Error()))
            return
        }

        // 假设创建用户成功，返回用户信息
        ctx.JSON(iris.StatusCreated, user)
    })
}

// main 函数是程序的入口点
func main() {
    // 创建一个新的App实例
    app := NewApp()

    // 设置路由和处理函数
    app.SetupRoutes()

    // 启动Iris框架的HTTP服务器
    // 监听8080端口，并处理进来的请求
    if err := app.Listen(":8080"); err != nil {
        panic(err)
    }
}

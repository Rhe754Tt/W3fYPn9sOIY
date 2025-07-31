// 代码生成时间: 2025-07-31 21:04:30
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// Handler 是一个处理HTTP请求的函数
// 它接收请求和响应对象，用于处理业务逻辑
func Handler(ctx iris.Context) {
    // 从请求中获取参数
    name := ctx.URLParam("name")

    // 检查参数是否为空
    if name == "" {
        // 如果参数为空，返回错误信息
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.WriteString("Parameter 'name' is required")
        return
    }

    // 业务逻辑处理
    fmt.Printf("Hello, %s!
", name)

    // 返回响应
    ctx.WriteString(fmt.Sprintf("Hello, %s!", name))
}

func main() {
    // 创建一个新的Iris应用程序
    app := iris.New()

    // 注册一个GET路由，使用Handler函数处理请求
    app.Get("/hello/{name}", Handler)

    // 启动服务器
    app.Listen(":8080")
}

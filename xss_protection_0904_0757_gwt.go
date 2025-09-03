// 代码生成时间: 2025-09-04 07:57:35
package main

import (
    "html"
    "log"
    "net/http"
    "github.com/kataras/iris/v12"
)

// XSSProtectionMiddleware 是中间件，用于防护XSS攻击
func XSSProtectionMiddleware(ctx iris.Context) {
    // 获取请求中的User-Agent头部
    userAgent := ctx.GetHeader("User-Agent")
    if len(userAgent) == 0 {
        // 如果User-Agent头部为空，禁止访问并返回错误
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.WriteString("Missing User-Agent header")
        return
    }
    // 对User-Agent进行HTML转义，以防止XSS攻击
    escapedUserAgent := html.EscapeString(userAgent)
    // 记录原始User-Agent和转义后的User-Agent
    log.Printf("Original User-Agent: %s
Escaped User-Agent: %s", userAgent, escapedUserAgent)
    ctx.Next()
}

func main() {
    app := iris.New()
    // 注册中间件
    app.Use(XSSProtectionMiddleware)

    // 定义一个简单的路由，用于测试中间件
    app.Get("/test", func(ctx iris.Context) {
        ctx.WriteString("Hello from XSS Protection Middleware")
    })

    // 启动服务器
    log.Println("Server is running at :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
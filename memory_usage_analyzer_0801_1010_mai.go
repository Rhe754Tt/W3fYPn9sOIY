// 代码生成时间: 2025-08-01 10:10:53
package main

import (
    "fmt"
    "os"
    "runtime"
    "time"
    "github.com/kataras/iris/v12"
)

// MemoryUsageAnalyzer 用于展示内存使用情况
func MemoryUsageAnalyzer() iris.HandlerFunc {
    return func(ctx iris.Context) {
        // 获取当前内存使用情况
        var m runtime.MemStats
        runtime.ReadMemStats(&m)

        // 计算内存使用百分比
        memUsage := float64(m.Alloc) / float64(m.Sys) * 100

        // 返回JSON格式的内存使用结果
        ctx.JSON(iris.StatusOK, map[string]float64{
            "memory_usage": memUsage,
        })
    }
}

func main() {
    app := iris.New()

    // 定义路由
    app.Get("/memory", MemoryUsageAnalyzer)

    // 设置端口并启动服务器
    port := 8080
    fmt.Printf("Server is running on port %d
", port)
    if err := app.Listen(":%d", port); err != nil {
        fmt.Printf("Error starting server: %s
", err.Error())
        os.Exit(1)
    }
}

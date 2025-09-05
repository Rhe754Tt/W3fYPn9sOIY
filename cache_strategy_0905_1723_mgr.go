// 代码生成时间: 2025-09-05 17:23:19
package main

import (
    "fmt"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/cache"
)

// CacheConfig 定义缓存配置
type CacheConfig struct {
    Duration time.Duration
}

// NewCacheConfig 创建一个新的缓存配置
func NewCacheConfig(duration time.Duration) *CacheConfig {
    return &CacheConfig{Duration: duration}
}

func main() {
    app := iris.New()
    // 使用默认的内存缓存
    cache := cache.New(cache.Config{
        adapter: cache.NewMemoryAdapter(),
    })
    cfg := NewCacheConfig(10 * time.Minute)

    // 设置缓存中间件
    app.Use(func(ctx iris.Context) {
        start := time.Now()
        ctx.Application().Logger().Infof("Start request: %s", ctx.Path())
        next := func() {
            ctx.Next()
        }
        ctx.Done(func(ctx iris.Context) {
            duration := time.Since(start)
            ctx.Application().Logger().Infof("End request: %s, Duration: %v", ctx.Path(), duration)
        })
        next()
    })

    // 定义路由和缓存策略
    app.Get("/data", func(ctx iris.Context) {
        // 尝试从缓存中获取数据
        if cache.Get(ctx, "data") == nil { {
            // 数据不在缓存中，获取数据并设置缓存
            data := fetchData()
            cache.Set(ctx, "data", data, cfg.Duration)
            fmt.Fprintf(ctx.ResponseWriter, "%s", data)
        } else {
            // 从缓存中获取数据
            if err := cache.Get(ctx, "data", nil); err == nil {
                fmt.Fprintf(ctx.ResponseWriter, "Data from cache")
            }
        }
    })

    // 启动服务器
    app.Listen(":8080")
}

// fetchData 模拟获取数据的函数
func fetchData() string {
    // 假设这里是从数据库或外部服务获取数据
    return "Cached Data"
}
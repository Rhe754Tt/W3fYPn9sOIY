// 代码生成时间: 2025-09-07 13:15:56
package main

import (
    "fmt"
    "math/rand"
    "time"
    "github.com/kataras/iris/v12"
)

// TestDataGenerator 是一个结构体，用于生成测试数据
type TestDataGenerator struct{}

// GenerateRandomString 生成一个随机字符串
func (g *TestDataGenerator) GenerateRandomString(length int) string {
    var result []rune
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    for i := 0; i < length; i++ {
        result = append(result, letters[r.Intn(len(letters))])
    }
    return string(result)
}

// GenerateRandomNumber 生成一个随机数字
func (g *TestDataGenerator) GenerateRandomNumber() int {
    return rand.Intn(100)
}

func main() {
    app := iris.New()

    // 初始化测试数据生成器
    dataGenerator := TestDataGenerator{}

    // 定义路由和处理函数
    app.Get("/test-data", func(ctx iris.Context) {
        // 生成测试数据
        randomString := dataGenerator.GenerateRandomString(10)
        randomNumber := dataGenerator.GenerateRandomNumber()

        // 返回测试数据
        ctx.JSON(iris.Map{
                "randomString": randomString,
                "randomNumber": randomNumber,
            })
    })

    // 启动服务
    app.Listen(":8080")
}

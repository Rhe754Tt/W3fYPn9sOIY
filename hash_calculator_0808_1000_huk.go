// 代码生成时间: 2025-08-08 10:00:25
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// HashCalculator 提供哈希值计算功能
type HashCalculator struct{}

// NewHashCalculator 创建一个新的 HashCalculator 实例
func NewHashCalculator() *HashCalculator {
    return &HashCalculator{}
}

// CalculateSHA256 计算输入字符串的 SHA-256 哈希值
func (h *HashCalculator) CalculateSHA256(input string) (string, error) {
    if input == "" {
        return "", fmt.Errorf("input cannot be empty")
    }

    // 计算 SHA-256哈希值
    hash := sha256.Sum256([]byte(input))

    // 将哈希值转换为十六进制字符串
    return hex.EncodeToString(hash[:]), nil
}

func main() {
    // 创建 Iris 应用
    app := iris.New()

    // 创建哈希值计算工具实例
    calculator := NewHashCalculator()

    // 定义路由，处理 POST 请求，接收字符串输入并计算哈希值
    app.Post("/hash", func(ctx iris.Context) {
        // 从请求体中获取字符串输入
        input := ctx.FormValue("input")

        // 验证输入是否为空
        if strings.TrimSpace(input) == "" {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "input cannot be empty",
            })
            return
        }

        // 计算哈希值
        sha256Hash, err := calculator.CalculateSHA256(input)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        // 返回计算结果
        ctx.JSON(iris.Map{
            "input": input,
            "hash": sha256Hash,
        })
    })

    // 设置应用在 8080 端口监听
    app.Listen(":8080")
}
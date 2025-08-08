// 代码生成时间: 2025-08-08 17:58:03
package main

import (
    "fmt"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// ApiResponse 结构体用于定义API响应的标准格式
type ApiResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// NewApiResponse 创建一个新的ApiResponse实例
func NewApiResponse(code int, message string, data interface{}) ApiResponse {
    return ApiResponse{
        Code:    code,
        Message: message,
        Data:    data,
    }
}

// ErrorResponse 创建一个错误响应
func ErrorResponse(code int, message string) ApiResponse {
    return NewApiResponse(code, message, nil)
}

// SuccessResponse 创建一个成功的响应
func SuccessResponse(data interface{}) ApiResponse {
    return NewApiResponse(http.StatusOK, "success", data)
}

func main() {
    // 创建一个Iris应用
    app := iris.New()

    // 定义一个路由，返回格式化的API响应
    app.Get("/api/response", func(ctx iris.Context) {
        // 模拟一些数据
        data := map[string]string{
            "key": "value",
        }

        // 创建成功的API响应
        response := SuccessResponse(data)

        // 将响应写入HTTP响应
        ctx.JSON(response)
    })

    // 定义一个路由，返回格式化的错误API响应
    app.Get("/api/error", func(ctx iris.Context) {
        // 创建错误的API响应
        response := ErrorResponse(http.StatusInternalServerError, "internal server error")

        // 将响应写入HTTP响应
        ctx.JSON(response)
    })

    // 启动Iris服务器
    app.Listen(":8080")
}

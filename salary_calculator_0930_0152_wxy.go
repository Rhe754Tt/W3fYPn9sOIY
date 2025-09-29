// 代码生成时间: 2025-09-30 01:52:20
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "net/http"
)

// SalaryCalculator 结构体，用于存储薪资计算器所需的数据
type SalaryCalculator struct {
    BaseSalary float64
    Bonus      float64
    Tax        float64
}

// CalculateSalary 方法，用于计算最终薪资
func (s *SalaryCalculator) CalculateSalary() float64 {
    return s.BaseSalary + s.Bonus - s.Tax
}

func main() {
    app := iris.New()

    // POST /calculateSalary 路径处理函数，用于接收薪资计算请求
    app.Post("/calculateSalary", func(ctx iris.Context) {
        decoder := ctx.JSON()
        var calculator SalaryCalculator
        // 错误处理，确保解码成功
        if err := decoder.Unmarshal(&calculator); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "Invalid input data",
            })
            return
        }
        // 计算薪资
        salary := calculator.CalculateSalary()
        // 返回计算结果
        ctx.JSON(iris.Map{
            "baseSalary": calculator.BaseSalary,
            "bonus": calculator.Bonus,
            "tax": calculator.Tax,
            "finalSalary": salary,
        })
    })

    // 启动服务
    fmt.Println("Server is running on :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}

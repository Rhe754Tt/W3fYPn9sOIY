// 代码生成时间: 2025-08-27 23:48:25
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/hero"
    "github.com/kataras/iris/v12/sessions"
)

// FormData 定义表单数据的结构
type FormData struct {
    Username string `json:"username" validate:"required,min=3,max=30"`
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"required,min=18,max=65"`
}

// Validate 实现表单验证的方法
func (f *FormData) Validate() error {
    if err := hero.Validate(f, nil); err != nil {
        return err
    }
    return nil
}

func main() {
    app := iris.New()
    // 设置视图和静态文件路径
    app.RegisterView(iris.HTML("./templates", ".html").Reload(true))
    app.HandleDir("/public", "./public")

    // Session配置
    sess := sessions.New(sessions.Config{
       Cookie:    "irissessionid",
      Expires:  15 * 60, // 15分钟过期
    })
    app.Use(sess.Load())

    // 注册表单路由
    app.Post("/form", func(ctx iris.Context) {
        var formData FormData
        // 解析JSON请求体到FormData结构
        if err := ctx.ReadJSON(&formData); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{"error": "Invalid JSON"})
            return
        }

        // 验证表单数据
        if err := formData.Validate(); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }

        // 存储表单数据到Session
        err := sess.Start(ctx)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{"error": "Session start failed"})
            return
        }
        sess.Set(ctx, "formData", formData)

        ctx.StatusCode(iris.StatusOK)
        ctx.JSON(iris.Map{"message": "Form data is valid and stored in session"})
    })

    // 启动Iris服务器
    app.Listen(":8080")
}
// 代码生成时间: 2025-08-04 14:27:35
package main

import (
    "fmt"
    "strings"

    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/sessions"
)

// FormValidator 结构体用于存储表单验证规则
type FormValidator struct {
    rules map[string]string
}

// NewFormValidator 创建一个新的表单验证器实例
func NewFormValidator() *FormValidator {
    return &FormValidator{
        rules: make(map[string]string),
    }
}

// AddRule 添加一个验证规则到验证器
func (fv *FormValidator) AddRule(field, rule string) {
    fv.rules[field] = rule
}

// Validate 验证表单数据
func (fv *FormValidator) Validate(form map[string]string) error {
    for field, rule := range fv.rules {
        if !ValidateField(form[field], rule) {
            return fmt.Errorf("invalid value for field '%s'", field)
        }
    }
    return nil
}

// ValidateField 根据规则验证单个字段
func ValidateField(value, rule string) bool {
    switch rule {
    case "required":
        return value != ""
    case "email":
        return strings.Contains(value, "@")
    case "minlength":
        // 这里只是一个简单的示例，实际应用中需要更复杂的逻辑
        minLength, _ := strconv.Atoi(strings.TrimPrefix(rule, "minlength:"))
        return len(value) >= minLength
    default:
        return true
    }
}

func main() {
    app := iris.New()
    // 设置会话
    sess := sessions.New(sessions.Config{Cookie: "iris-session"})
    app.Use(sess.Load())

    // 创建表单验证器
    validator := NewFormValidator()
    validator.AddRule("email", "required")
    validator.AddRule("email", "email")
    validator.AddRule("password", "minlength:8")

    app.Post("/form", func(ctx iris.Context) {
        form := map[string]string{
            "email": ctx.PostValue("email"),
            "password": ctx.PostValue("password"),
        }
        if err := validator.Validate(form); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }

        ctx.JSON(iris.Map{"message": "Form is valid"})
    })

    // 启动服务
    app.Listen(":8080")
}

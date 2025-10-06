// 代码生成时间: 2025-10-06 23:11:51
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/kataras/iris/v12"
)

// FirewallRule 定义了防火墙规则的结构
type FirewallRule struct {
    ID        uint   `json:"id"`
    RuleName  string `json:"rule_name"`
    Action    string `json:"action"`
    IPAddress string `json:"ip_address"`
}

// FirewallService 包含了管理防火墙规则的方法
type FirewallService struct {
    rules []FirewallRule
}

// NewFirewallService 创建一个新的防火墙服务实例
func NewFirewallService() *FirewallService {
    return &FirewallService{
        rules: []FirewallRule{},
    }
}

// AddRule 添加新的防火墙规则
func (fs *FirewallService) AddRule(rule FirewallRule) (uint, error) {
    fs.rules = append(fs.rules, rule)
    return rule.ID, nil
}

// GetRules 获取所有防火墙规则
func (fs *FirewallService) GetRules() []FirewallRule {
    return fs.rules
}

// DeleteRule 删除指定ID的防火墙规则
func (fs *FirewallService) DeleteRule(id uint) error {
    for i, rule := range fs.rules {
        if rule.ID == id {
            fs.rules = append(fs.rules[:i], fs.rules[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("rule with ID %d not found", id)
}

func main() {
    app := iris.New()
    service := NewFirewallService()

    // 定义API端点来管理防火墙规则
    api := app.Party("/api/firewall")

    // 添加新规则的端点
    api.Post("/add", func(ctx iris.Context) {
        var rule FirewallRule
        if err := ctx.ReadJSON(&rule); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{{"error": err.Error()}})
            return
        }
        rule.ID = uint(len(service.rules) + 1) // 简单的ID分配策略
        if _, err := service.AddRule(rule); err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.JSON(iris.Map{{"error": err.Error()}})
            return
        }
        ctx.StatusCode(http.StatusCreated)
        ctx.JSON(rule)
    })

    // 获取所有规则的端点
    api.Get("/rules", func(ctx iris.Context) {
        rules := service.GetRules()
        ctx.JSON(rules)
    })

    // 删除规则的端点
    api.Delete("/delete/{id:uint}", func(ctx iris.Context) {
        id := ctx.Params().GetUintDefault("id", 0)
        if err := service.DeleteRule(uint(id)); err != nil {
            ctx.StatusCode(http.StatusNotFound)
            ctx.JSON(iris.Map{{"error": err.Error()}})
            return
        }
        ctx.StatusCode(http.StatusOK)
        ctx.JSON(iris.Map{{"message": "rule deleted successfully"}})
    })

    // 启动Iris服务器
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("failed to start the server: %s", err)
    }
}
// 代码生成时间: 2025-08-23 12:43:54
package main

import (
    "fmt"
    "time"
    "github.com/kataras/iris/v12"
    "github.com/robfig/cron/v3"
)

// Scheduler 结构体用于封装定时任务调度器
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler 创建一个新的 Scheduler 实例
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(cron.WithSeconds()),
    }
}

// AddJob 添加一个新的定时任务到调度器
func (s *Scheduler) AddJob(spec string, cmd func()) (cron.EntryID, error) {
    entryID, err := s.cron.AddFunc(spec, cmd)
    if err != nil {
        return 0, err
    }
    return entryID, nil
}

// Start 启动定时任务调度器
func (s *Scheduler) Start() {
    s.cron.Start()
}

// Stop 停止定时任务调度器
func (s *Scheduler) Stop() {
    s.cron.Stop()
}

//定时任务执行的示例函数
func exampleTask() {
    fmt.Println("定时任务执行...")
}

func main() {
    // 创建 iris 应用
    app := iris.Default()

    // 创建定时任务调度器
    scheduler := NewScheduler()

    // 添加一个定时任务，每5秒执行一次
    entryID, err := scheduler.AddJob("*/5 * * * * *", exampleTask)
    if err != nil {
        fmt.Printf("添加定时任务失败：%v
", err)
        return
    }
    fmt.Printf("添加定时任务成功，任务ID：%d
", entryID)

    // 启动定时任务调度器
    scheduler.Start()

    // 启动 iris 服务器
    app.Get("/", func(ctx iris.Context) {
        ctx.WriteString("Hello, Scheduler!")
    })
    
    // 监听并服务 HTTP 请求
    if err := app.Listen(":8080"); err != nil {
        fmt.Printf("启动 Iris 服务失败：%v
", err)
    }
}

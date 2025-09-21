// 代码生成时间: 2025-09-21 23:26:54
package main

import (
    "fmt"
    "time"

    "github.com/kataras/iris/v12"
    "github.com/robfig/cron/v3"
)

// Scheduler 定时任务调度器
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler 创建一个新的调度器实例
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(),
    }
}

// Start 启动调度器
func (s *Scheduler) Start() {
    s.cron.Start()
}

// AddJob 添加一个定时任务
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    if _, err := s.cron.AddFunc(spec, cmd); err != nil {
        return err
    }
    return nil
}

func main() {
    // 创建IRIS应用程序
    app := iris.New()

    // 创建定时任务调度器实例
    scheduler := NewScheduler()

    // 启动调度器
    scheduler.Start()

    // 添加一个每5秒执行一次的定时任务
    if err := scheduler.AddJob("*/5 * * * *", func() {
        fmt.Println("定时任务执行中...")
    }); err != nil {
        fmt.Printf("添加定时任务失败：%v\
", err)
        return
    }

    // 启动IRIS服务器
    app.Listen(":8080")
}
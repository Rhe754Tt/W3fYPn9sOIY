// 代码生成时间: 2025-08-29 21:25:47
package main

import (
    "fmt"
    "time"
    "github.com/iris-contrib/middleware/cors"
    "github.com/kataras/iris/v12"
    "github.com/robfig/cron/v3"
)

// Scheduler 定时任务调度器
type Scheduler struct {
    Cron *cron.Cron
}

// NewScheduler 创建一个新的定时任务调度器
func NewScheduler() *Scheduler {
    return &Scheduler{
        Cron: cron.New(cron.WithSeconds()),
    }
}

// AddJob 添加一个新的定时任务
func (s *Scheduler) AddJob(spec string, cmd func()) error {
    _, err := s.Cron.AddFunc(spec, cmd)
    if err != nil {
        return err

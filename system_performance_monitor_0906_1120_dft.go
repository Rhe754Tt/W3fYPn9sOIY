// 代码生成时间: 2025-09-06 11:20:27
package main

import (
    "fmt"
    "os"
    "runtime"
    "strings"
    "time"

    "github.com/kardiwan/iris-metrics"
    "github.com/kevinburke/twilio-go"
)

// SystemPerformanceMonitor 结构体，用于存储系统性能监控相关参数
type SystemPerformanceMonitor struct {
    // metricsCollector 用于收集系统性能指标
    metricsCollector *irismetrics.Collector
    // twilioClient 用于发送短信通知
    twilioClient *twilio.Client
}

// NewSystemPerformanceMonitor 创建一个新的 SystemPerformanceMonitor 实例
func NewSystemPerformanceMonitor(accountSID, authToken, fromNumber string) *SystemPerformanceMonitor {
    collector := irismetrics.NewCollector()
    c := twilio.NewClient(accountSID, authToken, fromNumber)
    return &SystemPerformanceMonitor{
        metricsCollector: collector,
        twilioClient:     c,
    }
}

// Monitor 监控系统性能，如果发现异常则发送短信通知
func (spm *SystemPerformanceMonitor) Monitor(toNumber string, frequency time.Duration) {
    for {
        // 收集系统性能指标
        stats := spm.collectSystemMetrics()

        // 检查系统性能是否正常
        if !stats.IsHealthy() {
            // 发送短信通知
            spm.sendNotification(toNumber, stats)
        }

        // 等待下一个监控周期
        time.Sleep(frequency)
    }
}

// collectSystemMetrics 收集系统性能指标
func (spm *SystemPerformanceMonitor) collectSystemMetrics() *SystemMetrics {
    stats := NewSystemMetrics()

    // 收集内存使用情况
    stats.MemoryAllocated = runtime.MemStats.Alloc
    // 收集goroutine数量
    stats.GoRoutines = runtime.NumGoroutine()
    // 收集CPU使用情况（需要额外的库支持）
    // stats.CPUUsage = getCPUUsage()

    // 将系统性能指标添加到监控收集器
    spm.metricsCollector.Collect(stats)
    return stats
}

// sendNotification 发送短信通知
func (spm *SystemPerformanceMonitor) sendNotification(toNumber string, stats *SystemMetrics) {
    message := fmt.Sprintf("System performance alert: Memory: %d, GoRoutines: %d", stats.MemoryAllocated, stats.GoRoutines)
    _, err := spm.twilioClient.Messages.Create(toNumber, message)
    if err != nil {
        fmt.Printf("Error sending notification: %v", err)
    }
}

// SystemMetrics 结构体，用于存储系统性能指标
type SystemMetrics struct {
    MemoryAllocated uint64
    GoRoutines     int
    // CPUUsage       float64 // 需要额外的库支持
}

// IsHealthy 检查系统性能是否正常
func (sm *SystemMetrics) IsHealthy() bool {
    // 这里可以根据实际需求定义“健康”的标准
    // 例如，内存使用超过某个阈值或者goroutine数量超过某个值时，认为系统不正常
    return sm.MemoryAllocated < 1000000 && sm.GoRoutines < 100
}

func main() {
    // 创建 SystemPerformanceMonitor 实例
    spm := NewSystemPerformanceMonitor("YOUR_TWILIO_ACCOUNT_SID", "YOUR_TWILIO_AUTH_TOKEN", "YOUR_TWILIO_FROM_NUMBER")

    // 监控系统性能，每5分钟检查一次
    spm.Monitor("YOUR_TWILIO_TO_NUMBER", 5*time.Minute)
}

// 代码生成时间: 2025-08-09 05:18:02
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "runtime"
    "time"
    "github.com/kardianos/osext"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/net"
    "github.com/kataras/iris/v12"
)

// SystemMonitor contains system performance metrics
type SystemMonitor struct {
    Cpu        float64
    Mem        float64
    Disk       float64
    NetworkIn  float64
    NetworkOut float64
}

// GetSystemMetrics fetches various system metrics
func GetSystemMetrics() (*SystemMonitor, error) {
    var sm SystemMonitor
    var err error

    // Get CPU usage percentage
    sm.Cpu, err = cpu.Percent(0, false)
    if err != nil {
        return nil, err
    }

    // Get memory usage percentage
    sm.Mem, err = mem.VirtualMemory()
    if err != nil {
        return nil, err
    }
    sm.Mem = sm.Mem.UsedPercent

    // Get disk usage percentage
    sm.Disk, err = disk.Usage("/")
    if err != nil {
        return nil, err
    }
    sm.Disk = sm.Disk.UsedPercent

    // Get network I/O statistics (in bytes)
    stats, err := net.IOCounters(false)
    if err != nil {
        return nil, err
    }
    sm.NetworkIn = float64(stats[0].BytesRecv)
    sm.NetworkOut = float64(stats[0].BytesSent)

    return &sm, nil
}

// SystemMonitorHandler handles HTTP requests for system metrics
func SystemMonitorHandler(ctx iris.Context) {
    sm, err := GetSystemMetrics()
    if err != nil {
        ctx.StatusCode(http.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": fmt.Sprintf("Failed to get system metrics: %v", err),
        })
        return
    }

    ctx.JSON(iris.Map{
        "cpu":        sm.Cpu,
        "memory":     sm.Mem,
        "disk":       sm.Disk,
        "network_in": sm.NetworkIn,
        "network_out": sm.NetworkOut,
    })
}

func main() {
    app := iris.Default()

    // Register the system monitor handler
    app.Get("/system", SystemMonitorHandler)

    // Start the server
    port := "8080"
    fmt.Printf("System Monitor Service is running on port %s
", port)
    if err := app.Listen(":%s", iris.WithOptimizations()); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

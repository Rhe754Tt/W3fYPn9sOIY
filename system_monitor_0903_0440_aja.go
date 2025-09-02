// 代码生成时间: 2025-09-03 04:40:56
It fetches system metrics and exposes them via an HTTP endpoint.
*/

package main

import (
    "fmt"
    "os"
    "runtime"
    "syscall"
    "time"

    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/net"
    "github.com/kataras/iris/v12"
)

// SystemMetrics holds system metrics
type SystemMetrics struct {
    CpuUsage float64 `json:"cpu_usage"`
    MemoryUsage float64 `json:"memory_usage"`
    DiskUsage float64 `json:"disk_usage"`
    NetworkUsage float64 `json:"network_usage"`
}

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))

    // Define the route for system metrics
    app.Get("/metrics", func(ctx iris.Context) {
        metrics := fetchSystemMetrics()
        ctx.JSON(metrics)
    })

    // Start the IRIS server
    app.Listen(":8080")
}

// fetchSystemMetrics collects and returns system metrics
func fetchSystemMetrics() SystemMetrics {
    cpuPercent, _ := cpu.Percent(0, false)
    memory, _ := mem.VirtualMemory()
    disk, _ := disk.Usage("/")
    netIO, _ := net.IOCounters()

    return SystemMetrics{
        CpuUsage: cpuPercent[0],
        MemoryUsage: memory.UsedPercent,
        DiskUsage: disk.UsedPercent,
        NetworkUsage: float64(netIO.BytesSent+netIO.BytesRecv) / 1024 / 1024,
    }
}

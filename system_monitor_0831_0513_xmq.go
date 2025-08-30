// 代码生成时间: 2025-08-31 05:13:40
package main

import (
    "fmt"
    "net/http"
    "os"
    "regexp"
    "runtime"
    "time"

    "github.com/kardianos/osext"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/disk"
    "github.com/shirou/gopsutil/load"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/net"
)

// SystemMonitor is a struct containing system performance data
type SystemMonitor struct {
    // TODO: Add more fields as needed
    CpuUsage float64
    MemUsage float64
    DiskUsage float64
    NetStats  []net.IOCountersStat
    Load     load.AvgStat
}

// NewSystemMonitor creates a new SystemMonitor instance
func NewSystemMonitor() *SystemMonitor {
    return &SystemMonitor{}
}

// GetSystemInfo collects system performance data
func (sm *SystemMonitor) GetSystemInfo() error {
    // Get CPU usage
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return err
    }
    sm.CpuUsage = cpuPercent[0]

    // Get memory usage
    memStat, err := mem.VirtualMemory()
    if err != nil {
        return err
    }
    sm.MemUsage = memStat.UsedPercent

    // Get disk usage
    diskStat, err := disk.Usage("/")
    if err != nil {
        return err
    }
    sm.DiskUsage = diskStat.UsedPercent

    // Get network stats
    netStats, err := net.IOCounters(true)
    if err != nil {
        return err
    }
    sm.NetStats = netStats

    // Get load average
    loadStat, err := load.Avg()
    if err != nil {
        return err
    }
    sm.Load = loadStat

    return nil
}

// StartServer starts the HTTP server and listens for requests
func StartServer() {
    http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
        sm := NewSystemMonitor()
        if err := sm.GetSystemInfo(); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Respond with system performance data
        fmt.Fprintf(w, "CPU Usage: %.2f%%
", sm.CpuUsage)
        fmt.Fprintf(w, "Memory Usage: %.2f%%
", sm.MemUsage)
        fmt.Fprintf(w, "Disk Usage: %.2f%%
", sm.DiskUsage)
        fmt.Fprintf(w, "Network Stats: %+v
", sm.NetStats)
        fmt.Fprintf(w, "Load Averages: %+v
", sm.Load)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server started on port %s
", port)
    if err := http.ListenAndServe(":" + port, nil); err != nil {
        fmt.Printf("HTTP server failed: %v
", err)
        os.Exit(1)
    }
}

func main() {
    StartServer()
}
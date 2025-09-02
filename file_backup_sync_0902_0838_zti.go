// 代码生成时间: 2025-09-02 08:38:56
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"
)

const backupPath = "./backup/"

// SyncFiles 同步文件
func SyncFiles(sourcePath, destPath string) error {
    srcFiles, err := ioutil.ReadDir(sourcePath)
    if err != nil {
        return fmt.Errorf("读取源目录失败: %w", err)
    }

    for _, file := range srcFiles {
        srcFilePath := filepath.Join(sourcePath, file.Name())
        destFilePath := filepath.Join(destPath, file.Name())

        // 如果目标路径不存在，则创建
        if _, err := os.Stat(destFilePath); os.IsNotExist(err) {
            if err := os.MkdirAll(filepath.Dir(destFilePath), 0755); err != nil {
                return fmt.Errorf("创建目标目录失败: %w", err)
            }
        }

        // 复制文件
        if err := copyFile(srcFilePath, destFilePath); err != nil {
            return fmt.Errorf("复制文件失败: %w", err)
        }
    }

    return nil
}

// BackupFile 备份文件
func BackupFile(filePath, backupPath string) error {
    backupFilePath := filepath.Join(backupPath, filepath.Base(filePath))
    if _, err := os.Stat(backupFilePath); os.IsNotExist(err) {
        // 复制文件到备份目录
        if err := copyFile(filePath, backupFilePath); err != nil {
            return fmt.Errorf("备份文件失败: %w", err)
        }
    }
    return nil
}

// copyFile 复制文件
func copyFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("打开源文件失败: %w", err)
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("创建目标文件失败: %w", err)
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        return fmt.Errorf("复制文件失败: %w", err)
    }

    return nil
}

// HandleBackup 处理备份请求
func HandleBackup(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    filePath := r.URL.Query().Get("file")
    if filePath == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if err := BackupFile(filePath, backupPath); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "备份文件失败: %s", err)
        return
    }

    fmt.Fprintln(w, "文件备份成功")
}

// HandleSync 处理同步请求
func HandleSync(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    sourcePath := r.URL.Query().Get("source")
    destPath := r.URL.Query().Get("dest")
    if sourcePath == "" || destPath == "" {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    if err := SyncFiles(sourcePath, destPath); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "同步文件失败: %s", err)
        return
    }

    fmt.Fprintln(w, "文件同步成功")
}

func main() {
    // 确保备份目录存在
    if _, err := os.Stat(backupPath); os.IsNotExist(err) {
        if err := os.MkdirAll(backupPath, 0755); err != nil {
            log.Fatalf("创建备份目录失败: %s", err)
        }
    }

    app := iris.Default()

    // 注册备份和同步路由
    app.Post("/backup", HandleBackup)
    app.Post("/sync", HandleSync)

    // 启动服务
    log.Printf("服务启动成功，监听端口: %d", 8080)
    if err := app.Listen(":8080"); err != nil {
        log.Fatalf("服务启动失败: %s", err)
    }
}
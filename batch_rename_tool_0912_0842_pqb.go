// 代码生成时间: 2025-09-12 08:42:45
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

// Renamer 结构体包含重命名函数的配置信息
type Renamer struct {
    // 源目录
    SourceDir string
    // 目标目录
    TargetDir string
    // 重命名规则
    RenameRule func(string) string
}

// NewRenamer 返回一个配置好的 Renamer 实例
func NewRenamer(sourceDir, targetDir string, renameRule func(string) string) *Renamer {
    return &Renamer{
        SourceDir: sourceDir,
        TargetDir: targetDir,
        RenameRule: renameRule,
    }
}

// Rename 遍历源目录中的所有文件，并根据规则重命名文件
func (r *Renamer) Rename() error {
    if _, err := os.Stat(r.SourceDir); os.IsNotExist(err) {
        return fmt.Errorf("源目录不存在: %s", r.SourceDir)
    }

    err := filepath.Walk(r.SourceDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            targetName := r.RenameRule(info.Name())
            targetPath := filepath.Join(r.TargetDir, targetName)

            // 检查目标文件是否已存在
            if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
                return fmt.Errorf("目标文件已存在: %s", targetPath)
            }

            // 重命名文件
            if err := os.Rename(path, targetPath); err != nil {
                return fmt.Errorf("重命名文件失败: %s -> %s, 错误: %v", path, targetPath, err)
            }

            fmt.Printf("文件已重命名: %s -> %s
", path, targetPath)
        }
        return nil
    })
    return err
}

func main() {
    // 设置IRIS服务器
    app := iris.New()
    app.Get("/rename", func(ctx iris.Context) {
        // 调用重命名函数
        renamer := NewRenamer("./source", "./target", func(filename string) string {
            // 定义重命名规则
            nameWithoutExt := strings.TrimSuffix(filename, filepath.Ext(filename))
            timestamp := time.Now().Format("20060102-150405")
            return fmt.Sprintf("%s_%s%s", nameWithoutExt, timestamp, filepath.Ext(filename))
        })
        if err := renamer.Rename(); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
            return
        }
        ctx.WriteString("文件重命名成功")
    })
    app.Listen(":8080")
}

// 代码生成时间: 2025-09-17 12:12:32
package main

import (
    "archive/zip"
    "io"
    "io/fs"
    "log"
    "os"
    "path/filepath"
)

// Unzip解压文件到指定目录
func Unzip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    // 遍历压缩包中的文件
    for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        // 创建目标文件
        destPath := filepath.Join(dest, f.Name)
        
        // 确保目录存在
        if f.FileInfo().IsDir() {
            os.MkdirAll(destPath, fs.FileMode(f.Mode()))
            continue
        }

        fDir := filepath.Dir(destPath)
        if _, err := os.Stat(fDir); err != nil && !os.IsNotExist(err) {
            return err
        }
        if err := os.MkdirAll(fDir, fs.FileMode(f.Mode())); err != nil {
            return err
        }

        dw, err := os.Create(destPath)
        if err != nil {
            return err
        }
        defer dw.Close()

        _, err = io.Copy(dw, rc)
        if err != nil {
            return err
        }
        dw.Close()

        // 保持文件权限
        if err := os.Chmod(destPath, fs.FileMode(f.Mode())); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    // 示例：解压文件到指定目录
    err := Unzip("example.zip", "extracted/")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Unzipped successfully!")
}
// 代码生成时间: 2025-08-19 14:48:09
package main

import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/kataras/iris/v12"
)

// BackupData represents the data structure for backup data
type BackupData struct {
    Data  []byte
    Hash  string
    Name  string
    Error string
}

// backupFile creates a backup of the specified file
func backupFile(ctx iris.Context, filePath string) (*BackupData, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    data, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

    // Create a hash for the backup data
    hash := fmt.Sprintf("%x", md5.Sum(data))

    // Save the backup data to a new file
    backupPath := fmt.Sprintf("%s.bak", filePath)
    err = ioutil.WriteFile(backupPath, data, 0644)
    if err != nil {
        return nil, err
    }

    return &BackupData{Data: data, Hash: hash, Name: backupPath}, nil
}

// restoreFile restores the backup file to its original
func restoreFile(ctx iris.Context, filePath, backupPath string) error {
    backupFile, err := os.Open(backupPath)
    if err != nil {
        return err
    }
    defer backupFile.Close()

    backupData, err := ioutil.ReadAll(backupFile)
    if err != nil {
        return err
    }

    // Write the backup data to the original file
    return ioutil.WriteFile(filePath, backupData, 0644)
}

func main() {
    app := iris.New()

    // Endpoint for backing up a file
    app.Post("/backup", func(ctx iris.Context) {
        filePath := ctx.URLParam("file")
        if filePath == "" {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "File path is required",
            })
            return
        }

        backupData, err := backupFile(ctx, filePath)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        ctx.JSON(backupData)
    })

    // Endpoint for restoring a file from its backup
    app.Post("/restore", func(ctx iris.Context) {
        filePath := ctx.URLParam("file")
        backupPath := ctx.URLParam("backup")
        if filePath == "" || backupPath == "" {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "Both file and backup paths are required",
            })
            return
        }

        err := restoreFile(ctx, filePath, backupPath)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        ctx.JSON(iris.Map{
            "message": "File restored successfully",
        })
    })

    // Start the IRIS server
    log.Fatal(app.Listen(":8080"))
}

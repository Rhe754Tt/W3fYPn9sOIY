// 代码生成时间: 2025-09-18 00:42:12
Features:
- Code structure is clear and understandable.
- Contains appropriate error handling.
- Includes necessary comments and documentation.
- Follows GOLANG best practices.
- Ensures code maintainability and extensibility.
*/

package main

import (
    "bytes"
    "compress/zip"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/kataras/iris/v12" // iris framework
)

// unzipPath is the directory where files will be extracted.
var unzipPath string

// Setup sets up the necessary environment variables and configurations.
func Setup() {
    unzipPath = "./extracted_files" // default extraction path
    if _, err := os.Stat(unzipPath); os.IsNotExist(err) {
        os.MkdirAll(unzipPath, os.ModePerm)
    }
}

// Unzip extracts the contents of a zip file into the specified path.
func Unzip(zipFile string) error {
    reader, err := zip.OpenReader(zipFile)
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        filePath := filepath.Join(unzipPath, file.Name)
        if file.FileInfo().IsDir() {
            // Create directory.
            os.MkdirAll(filePath, os.ModePerm)
            continue
        }

        // Create file.
        fileReader, err := file.Open()
        if err != nil {
            return err
        }
        defer fileReader.Close()

        targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
        if err != nil {
            return err
        }
        defer targetFile.Close()

        _, err = io.Copy(targetFile, fileReader)
        if err != nil {
            return err
        }
    }
    return nil
}

// HandleUnzipRequest handles the HTTP request to unzip a file.
func HandleUnzipRequest(ctx iris.Context) {
    file, _, err := ctx.FormFile("file")
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("Error: Failed to retrieve file.")
        return
    }
    defer file.Close()

    tempFile, err := ioutil.TempFile("", "zipfile-")
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("Error: Failed to create temporary file.")
        return
    }
    defer tempFile.Close()
    defer os.Remove(tempFile.Name())

    _, err = io.Copy(tempFile, file)
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("Error: Failed to copy file.")
        return
    }

    err = Unzip(tempFile.Name())
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString(fmt.Sprintf("Error: Failed to unzip file - %s", err.Error()))
        return
    }

    ctx.StatusCode(iris.StatusOK)
    ctx.WriteString("File successfully unzipped.")
}

func main() {
    Setup()
    app := iris.New()
    app.Post("/unzip", HandleUnzipRequest)
    app.Listen(":8080")
}
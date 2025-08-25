// 代码生成时间: 2025-08-26 02:31:35
package main

import (
    "archive/zip"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Unzipper is a struct that contains the destination path for the unzip operation.
type Unzipper struct {
    Destination string
}

// NewUnzipper creates a new Unzipper instance with a specified destination.
func NewUnzipper(destination string) *Unzipper {
    return &Unzipper{Destination: destination}
}

// Unzip takes a zip file path and decompresses it to the destination.
func (u *Unzipper) Unzip(zipFilePath string) error {
    // Open the zip file
    file, err := os.Open(zipFilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    // Create a reader from the zip file
    reader, err := zip.OpenReader(zipFilePath)
    if err != nil {
        return err
    }
    defer reader.Close()

    // Iterate through the files in the zip.
    for _, zipFile := range reader.File {
        // Get the file information and create the file path
        fPath := filepath.Join(u.Destination, zipFile.Name)
        if zipFile.FileInfo().IsDir() {
            // Create directory if needed
            os.MkdirAll(fPath, os.ModePerm)
            continue
        }

        // Create the file
        if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
            return err
        }
        fileToWrite, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
        if err != nil {
            return err
        }
        defer fileToWrite.Close()

        // Copy the contents of the zip file to the new file
        fileInZip, err := zipFile.Open()
        if err != nil {
            return err
        }
        defer fileInZip.Close()
        _, err = io.Copy(fileToWrite, fileInZip)
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    zipFilePath := flag.String("zip", "", "Path to the zip file to unzip")
    destPath := flag.String("dest", "./", "Destination directory to unzip the files")
    flag.Parse()

    if *zipFilePath == "" {
        log.Fatal("No zip file path provided")
    }

    unzipper := NewUnzipper(*destPath)
    if err := unzipper.Unzip(*zipFilePath); err != nil {
        fmt.Printf("Error unzipping file: %s
", err)
    } else {
        fmt.Printf("Unzipped successfully to %s
", *destPath)
    }
}

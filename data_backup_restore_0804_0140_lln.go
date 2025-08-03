// 代码生成时间: 2025-08-04 01:40:36
package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// Constants for backup file naming and path
const (
    backupDirName     = "backups"
    backupFileNameFmt = "data_backup_%s.tar.gz"
)

// PerformBackup takes a data directory and creates a backup of it.
func PerformBackup(dataDir string) error {
    // Create a timestamp for the backup file name
    timestamp := time.Now().Format("20060102_150405")
    backupFileName := fmt.Sprintf(backupFileNameFmt, timestamp)

    // Create the backup directory if it doesn't exist
    err := os.MkdirAll(backupDirName, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create backup directory: %v", err)
    }

    // Full path for the backup file
    backupFilePath := filepath.Join(backupDirName, backupFileName)

    // Create a new tar.gz file for backup
    outFile, err := os.Create(backupFilePath)
    if err != nil {
        return fmt.Errorf("failed to create backup file: %v", err)
    }
    defer outFile.Close()

    // Create a gzip writer for compression
    gzipWriter := gzip.NewWriter(outFile)
    defer gzipWriter.Close()

    // Create a tar writer for archiving
    tarWriter := tar.NewWriter(gzipWriter)
    defer tarWriter.Close()

    // Walk through the data directory and add files to the tar
    err = filepath.WalkDir(dataDir, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // Ignore the backup directory itself
        if d.IsDir() && path == backupDirName {
            return filepath.SkipDir
        }

        // Open the file
        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        // Get file information
        stat, err := d.Info()
        if err != nil {
            return err
        }

        // Create a tar header
        tarHeader, err := tar.FileInfoHeader(stat, path)
        if err != nil {
            return err
        }
        tarHeader.Name = strings.TrimPrefix(path, dataDir+"/") // Remove dataDir from the path

        // Write header to tar
        if err = tarWriter.WriteHeader(tarHeader); err != nil {
            return err
        }

        // Write file content to tar
        if !stat.IsDir() {
            _, err = io.Copy(tarWriter, file)
            if err != nil {
                return err
            }
        }

        return nil
    })

    if err != nil {
        return fmt.Errorf("backup failed: %v", err)
    }

    return nil
}

// PerformRestore takes a backup file and restores the data.
func PerformRestore(backupFilePath string) error {
    // Open the backup file
    backupFile, err := os.Open(backupFilePath)
    if err != nil {
        return fmt.Errorf("failed to open backup file: %v", err)
    }
    defer backupFile.Close()

    // Create a gzip reader for decompression
    gzipReader, err := gzip.NewReader(backupFile)
    if err != nil {
        return fmt.Errorf("failed to create gzip reader: %v", err)
    }
    defer gzipReader.Close()

    // Create a tar reader for extracting
    tarReader := tar.NewReader(gzipReader)

    // Extract files from the tar
    for {
        header, err := tarReader.Next()
        if err == io.EOF {
            break // End of archive
        }
        if err != nil {
            return fmt.Errorf("failed to read tar header: %v", err)
        }

        // Skip directories
        if header.Typeflag == tar.TypeDir {
            continue
        }

        // Create the directory structure
        dirPath := filepath.Dir(header.Name)
        if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
            return fmt.Errorf("failed to create directory %s: %v", dirPath, err)
        }

        // Create the file
        outFile, err := os.Create(header.Name)
        if err != nil {
            return fmt.Errorf("failed to create file %s: %v", header.Name, err)
        }
        defer outFile.Close()

        // Copy the file content
        if _, err = io.Copy(outFile, tarReader); err != nil {
            return fmt.Errorf("failed to copy file content: %v", err)
        }
    }

    return nil
}

func main() {
    dataDir := "./data" // The directory containing the data to backup
    backupResult := PerformBackup(dataDir)
    if backupResult != nil {
        log.Fatalf("Backup failed: %v", backupResult)
    } else {
        fmt.Println("Backup successful!")
    }

    // Specify the path to the backup file you want to restore
    backupFilePath := "./backups/data_backup_20231122_091500.tar.gz"
    restoreResult := PerformRestore(backupFilePath)
    if restoreResult != nil {
        log.Fatalf("Restore failed: %v", restoreResult)
    } else {
        fmt.Println("Restore successful!")
    }
}

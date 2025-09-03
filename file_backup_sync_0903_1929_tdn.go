// 代码生成时间: 2025-09-03 19:29:10
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "time"
)

// BackupSyncer defines the structure for the backup and sync utility
type BackupSyncer struct {
    SourcePath    string
    DestinationPath string
    BackupTime    time.Time
}

// NewBackupSyncer creates a new instance of BackupSyncer
func NewBackupSyncer(sourcePath, destinationPath string) *BackupSyncer {
    return &BackupSyncer{
        SourcePath:    sourcePath,
        DestinationPath: destinationPath,
        BackupTime:    time.Now(),
    }
}

// Sync syncs the files from the source to the destination path
func (b *BackupSyncer) Sync() error {
    // Check if the source path exists
    if _, err := os.Stat(b.SourcePath); os.IsNotExist(err) {
        return fmt.Errorf("source path does not exist: %w", err)
    }

    // Create destination directory if it doesn't exist
    if err := os.MkdirAll(b.DestinationPath, 0755); err != nil {
        return fmt.Errorf("failed to create destination directory: %w", err)
    }

    // Walk through the source directory and copy files to the destination
    return filepath.Walk(b.SourcePath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error walking through source path: %w", err)
        }

        if info.IsDir() {
            return nil
        }

        // Construct destination file path and create file
        destinationFile, err := os.Create(filepath.Join(b.DestinationPath, filepath.Base(path)))
        if err != nil {
            return fmt.Errorf("failed to create destination file: %w", err)
        }
        defer destinationFile.Close()

        // Open source file and copy its content to the destination file
        sourceFile, err := os.Open(path)
        if err != nil {
            return fmt.Errorf("failed to open source file: %w", err)
        }
        defer sourceFile.Close()

        if _, err := io.Copy(destinationFile, sourceFile); err != nil {
            return fmt.Errorf("failed to copy file content: %w", err)
        }

        return nil
    })
}

// Backup creates a backup of the source directory at the destination path
func (b *BackupSyncer) Backup() error {
    // Check if the source path exists
    if _, err := os.Stat(b.SourcePath); os.IsNotExist(err) {
        return fmt.Errorf("source path does not exist: %w", err)
    }

    // Create backup directory if it doesn't exist
    backupDir := filepath.Join(b.DestinationPath, fmt.Sprintf("backup_%s", b.BackupTime.Format("2006_01_02_150405")))
    if err := os.MkdirAll(backupDir, 0755); err != nil {
        return fmt.Errorf("failed to create backup directory: %w", err)
    }

    // Copy the entire source directory to the backup directory
    return CopyDirectory(b.SourcePath, backupDir)
}

// CopyDirectory copies the entire directory recursively
func CopyDirectory(src, dst string) error {
    return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return fmt.Errorf("error walking through source path: %w", err)
        }

        if info.IsDir() {
            // Create sub-directory in destination path
            return os.MkdirAll(filepath.Join(dst, filepath.Base(path)), info.Mode())
        }

        // Create destination file path and create file
        destinationFile, err := os.Create(filepath.Join(dst, filepath.Base(path)))
        if err != nil {
            return fmt.Errorf("failed to create destination file: %w", err)
        }
        defer destinationFile.Close()

        // Open source file and copy its content to the destination file
        sourceFile, err := os.Open(path)
        if err != nil {
            return fmt.Errorf("failed to open source file: %w", err)
        }
        defer sourceFile.Close()

        if _, err := io.Copy(destinationFile, sourceFile); err != nil {
            return fmt.Errorf("failed to copy file content: %w", err)
        }

        return nil
    })
}

func main() {
    sourcePath := "./source"
    destinationPath := "./destination"

    backupSyncer := NewBackupSyncer(sourcePath, destinationPath)

    // Perform backup
    if err := backupSyncer.Backup(); err != nil {
        log.Fatalf("failed to backup: %v", err)
    }

    // Perform sync
    if err := backupSyncer.Sync(); err != nil {
        log.Fatalf("failed to sync: %v", err)
    }

    fmt.Println("Backup and sync operations completed successfully.")
}

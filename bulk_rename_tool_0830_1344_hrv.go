// 代码生成时间: 2025-08-30 13:44:48
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

// BulkRenamer contains the base directory path and a function to rename files.
type BulkRenamer struct {
    BaseDir string
}

// NewBulkRenamer creates a new BulkRenamer instance.
func NewBulkRenamer(baseDir string) *BulkRenamer {
    return &BulkRenamer{
        BaseDir: baseDir,
    }
}

// RenameFiles renames all files in the base directory according to a specified pattern.
func (br *BulkRenamer) RenameFiles(pattern string) error {
    files, err := ioutil.ReadDir(br.BaseDir)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue // Skip directories.
        }

        oldPath := filepath.Join(br.BaseDir, file.Name())
        newPath := filepath.Join(br.BaseDir, generateNewName(pattern, file.Name()))

        if err := os.Rename(oldPath, newPath); err != nil {
            return fmt.Errorf("failed to rename file %s to %s: %w", oldPath, newPath, err)
        }
        fmt.Printf("Renamed %s to %s
", oldPath, newPath)
    }
    return nil
}

// generateNewName creates a new file name based on the given pattern and original file name.
func generateNewName(pattern string, originalName string) string {
    nameParts := strings.SplitN(originalName, ".", 2)
    extension := nameParts[1]
    baseName := fmt.Sprintf("%s_%d.%s", pattern, time.Now().Unix(), extension)
    return baseName
}

func main() {
    // Initialize the BulkRenamer with the base directory path.
    br := NewBulkRenamer("./")
    // Define the pattern for renaming files, e.g., adding a timestamp.
    pattern := "renamed"

    // Call the RenameFiles method to perform the renaming.
    if err := br.RenameFiles(pattern); err != nil {
        log.Fatalf("An error occurred during file renaming: %s
", err)
    }
}

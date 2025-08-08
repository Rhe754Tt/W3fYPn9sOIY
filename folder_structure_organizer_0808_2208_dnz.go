// 代码生成时间: 2025-08-08 22:08:32
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

// FolderStructureOrganizer represents the main struct for organizing folder structure.
type FolderStructureOrganizer struct {
    BasePath string
}

// NewFolderStructureOrganizer creates a new FolderStructureOrganizer instance.
func NewFolderStructureOrganizer(basePath string) *FolderStructureOrganizer {
    return &FolderStructureOrganizer{
        BasePath: basePath,
    }
}

// Organize takes care of organizing the folder structure.
func (fso *FolderStructureOrganizer) Organize() error {
    // Get all files and directories from the base path.
    files, err := ioutil.ReadDir(fso.BasePath)
    if err != nil {
        return fmt.Errorf("failed to read base path directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            // Skip if it's a directory.
            continue
        }

        // Construct file path.
        filePath := filepath.Join(fso.BasePath, file.Name())

        // Extract file extension.
        extension := strings.TrimPrefix(filepath.Ext(filePath), ".")

        // Create directory based on file extension if it doesn't exist.
        if _, err := os.Stat(filepath.Join(fso.BasePath, extension)); os.IsNotExist(err) {
            if err := os.MkdirAll(filepath.Join(fso.BasePath, extension), 0755); err != nil {
                return fmt.Errorf("failed to create directory: %w", err)
            }
        }

        // Move file to its respective directory based on file extension.
        targetPath := filepath.Join(fso.BasePath, extension, file.Name())
        if err := os.Rename(filePath, targetPath); err != nil {
            return fmt.Errorf("failed to move file: %w", err)
        }
    }

    return nil
}

func main() {
    // Create a new FolderStructureOrganizer with the specified base path.
    organizer := NewFolderStructureOrganizer("./")

    // Organize the folder structure.
    if err := organizer.Organize(); err != nil {
        fmt.Printf("An error occurred: %s
", err)
    } else {
        fmt.Println("Folder structure organized successfully.")
    }
}
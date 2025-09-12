// 代码生成时间: 2025-09-12 21:32:59
package main

import (
    "fmt"
    "image"
    "image/jpeg"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"
    "iris/v12"
)

// ImageResizer struct contains configuration for resizing images.
type ImageResizer struct {
    Width, Height int
    Quality        int
}

// NewImageResizer creates a new ImageResizer instance with default values.
func NewImageResizer(width, height int, quality int) *ImageResizer {
    return &ImageResizer{
        Width:  width,
        Height: height,
        Quality: quality,
    }
}

// ResizeImage resizes an image to the specified dimensions.
func (r *ImageResizer) ResizeImage(filePath string) error {
    srcFile, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open image file: %w", err)
    }
    defer srcFile.Close()

    img, _, err := image.Decode(srcFile)
    if err != nil {
        return fmt.Errorf("failed to decode image: %w", err)
    }

    resizedImg := image.NewRGBA(image.Rect(0, 0, r.Width, r.Height))
    resizedImg = resize.Resize(uint(r.Width), uint(r.Height), img, resize.Lanczos3)

    resizedFile, err := os.Create(filePath + "_resized.jpg")
    if err != nil {
        return fmt.Errorf("failed to create resized image file: %w", err)
    }
    defer resizedFile.Close()

    err = jpeg.Encode(resizedFile, resizedImg, &jpeg.Options{Quality: r.Quality})
    if err != nil {
        return fmt.Errorf("failed to encode resized image: %w", err)
    }

    return nil
}

// BatchResize resizes all images in a directory to the specified dimensions.
func (r *ImageResizer) BatchResize(directory string) error {
    files, err := ioutil.ReadDir(directory)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() && filepath.Ext(file.Name()) == ".jpg" {
            filePath := filepath.Join(directory, file.Name())
            if err := r.ResizeImage(filePath); err != nil {
                fmt.Printf("Error resizing image %s: %s
", file.Name(), err)
                continue
            }
        }
    }

    return nil
}

func main() {
    app := iris.New()
    // Configuration for the image resizer.
    resizer := NewImageResizer(800, 600, 75)

    app.Post("/resize", func(ctx iris.Context) {
        directory := ctx.URLParam("directory")
        if directory == "" {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.WriteString("Directory parameter is required.")
            return
        }

        if err := resizer.BatchResize(directory); err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString("Error resizing images: " + err.Error())
            return
        }

        ctx.WriteString("Images have been resized successfully.")
    })

    // Start the IRIS server.
    app.Listen(":8080")
}

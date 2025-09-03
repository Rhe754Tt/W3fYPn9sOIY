// 代码生成时间: 2025-09-03 15:31:23
package main

import (
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "github.com/kataras/iris/v12"
    "github.com/nfnt/resize"
)

// ImageResizer 是一个处理图片尺寸调整的结构
type ImageResizer struct {
    targetWidth, targetHeight int
}

// NewImageResizer 创建并初始化一个ImageResizer对象
func NewImageResizer(width, height int) *ImageResizer {
    return &ImageResizer{
        targetWidth:  width,
        targetHeight: height,
    }
}

// ResizeImage 调整给定图片的尺寸
func (r *ImageResizer) ResizeImage(srcPath, destPath string) error {
    srcFile, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    srcImg, _, err := image.Decode(srcFile)
    if err != nil {
        return err
    }

    destImg := resize.Resize(uint(r.targetWidth), uint(r.targetHeight), srcImg, resize.Lanczos3)

    destFile, err := os.Create(destPath)
    if err != nil {
        return err
    }
    defer destFile.Close()

    if destPath[len(destPath)-3:] == "jpg" || destPath[len(destPath)-3:] == "JPG" {
        err = jpeg.Encode(destFile, destImg, nil)
    } else if destPath[len(destPath)-3:] == "png" || destPath[len(destPath)-3:] == "PNG" {
        err = png.Encode(destFile, destImg)
    }
    if err != nil {
        return err
    }

    return nil
}

func main() {
    app := iris.New()
    app.Handle("GET", "/resize", func(ctx iris.Context) {
        // 这里可以添加更多的逻辑来处理请求参数，例如图片路径和目标尺寸
        // 例如：
        // srcPath := ctx.URLParam("src")
        // destPath := ctx.URLParam("dest")
        // width, _ := ctx.URLParamInt("width")
        // height, _ := ctx.URLParamInt("height")
        // resizer := NewImageResizer(width, height)
        // err := resizer.ResizeImage(srcPath, destPath)
        // if err != nil {
        //     ctx.StatusCode(iris.StatusInternalServerError)
        //     ctx.Writef("Error resizing image: %v", err)
        //     return
        // }
        // ctx.StatusCode(iris.StatusNoContent)
    })

    app.Listen(":8080")
}

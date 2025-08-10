// 代码生成时间: 2025-08-10 08:22:24
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"

    "github.com/kataras/iris/v12"
)

// ScrapeContent 定义了一个网页抓取的结构体，包含网页内容
type ScrapeContent struct {
    Content string `json:"content"`
}

// scrapeWebsite 函数用于从给定的URL抓取网页内容
func scrapeWebsite(url string) (*ScrapeContent, error) {
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }
    content := strings.TrimSpace(string(body))
    return &ScrapeContent{Content: content}, nil
}

func main() {
    app := iris.New()

    // 设置路由，用于抓取网页内容
    app.Get("/scrape/{url}", func(ctx iris.Context) {
        url := ctx.Params().Get("url")
        result, err := scrapeWebsite(url)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": fmt.Sprintf("Failed to scrape website: %s", err),
            })
            return
        }
        ctx.JSON(result)
    })

    // 设置服务器监听的端口和超时时间
    app.Listen(":8080", iris.WithServerTimeout(5 * time.Second))
}

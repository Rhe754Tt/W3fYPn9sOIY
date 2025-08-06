// 代码生成时间: 2025-08-06 22:42:37
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// WebContentScraper 用于定义抓取网页内容的结构
type WebContentScraper struct {
    client *http.Client
}

// NewWebContentScraper 创建一个新的 WebContentScraper 实例
func NewWebContentScraper() *WebContentScraper {
    // 设置超时时间为 5 秒
    return &WebContentScraper{
        client: &http.Client{
            Timeout: 5 * time.Second,
        },
    }
}

// FetchContent 从给定的网址抓取内容
func (s *WebContentScraper) FetchContent(url string) (string, error) {
    resp, err := s.client.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // 检查状态码是否为 200 OK
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to fetch content: status code %d", resp.StatusCode)
    }

    // 读取响应体的内容
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    // 返回内容的字符串表示
    return string(content), nil
}

func main() {
    // 创建一个新的 WebContentScraper 实例
    scraper := NewWebContentScraper()

    // Iris 设置
    app := iris.New()
    app.Get("/scraper/{url}", func(ctx iris.Context) {
        // 从 URL 参数中获取网址
        url := ctx.Params().Get("url")

        // 抓取网页内容
        content, err := scraper.FetchContent(url)
        if err != nil {
            // 如果发生错误，返回错误信息
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.WriteString(err.Error())
            return
        }

        // 返回抓取到的内容
        ctx.WriteString(content)
    })

    // 启动 Iris 服务器
    app.Listen(":8080")
}

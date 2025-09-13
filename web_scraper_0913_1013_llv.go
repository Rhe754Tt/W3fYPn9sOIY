// 代码生成时间: 2025-09-13 10:13:30
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "log"
    "golang.org/x/net/html"
)

// fetchPage 获取网页内容
func fetchPage(url string) (string, error) {
    // 发送HTTP GET请求
    response, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer response.Body.Close()

    // 读取响应体
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", err
    }

    // 将响应体转换为字符串
    return string(body), nil
}

// extractText 提取HTML中的文本
func extractText(n *html.Node) string {
    if n.Type == html.TextNode {
        return strings.TrimSpace(n.Data)
    }
    var text strings.Builder
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        text.WriteString(extractText(c))
    }
    return text.String()
}

// scrapeWebsite 抓取网页内容并提取文本
func scrapeWebsite(url string) (string, error) {
    // 获取网页内容
    pageContent, err := fetchPage(url)
    if err != nil {
        return "", err
    }

    // 解析HTML
    doc, err := html.Parse(strings.NewReader(pageContent))
    if err != nil {
        return "", err
    }

    // 提取文本
    return extractText(doc), nil
}

func main() {
    // 要抓取的网页URL
    url := "http://example.com"

    // 抓取网页内容
    text, err := scrapeWebsite(url)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(text)
}

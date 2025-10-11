// 代码生成时间: 2025-10-11 19:59:25
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/kataras/iris/v12"
)

// Price represents the structure of the price data.
type Price struct {
    ProductID string  "json:"product_id"
    Price     float64 "json:"price"
    Timestamp int64   "json:"timestamp"
}

// MonitorService is the service responsible for monitoring prices.
type MonitorService struct {
    client *http.Client
}

// NewMonitorService creates a new MonitorService instance.
func NewMonitorService() *MonitorService {
    return &MonitorService{
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

// FetchPrice retrieves the price from the given URL, simulating a price check.
func (s *MonitorService) FetchPrice(url string) (*Price, error) {
    resp, err := s.client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("non-200 response: %s", resp.Status)
    }
    var price Price
    if err := json.NewDecoder(resp.Body).Decode(&price); err != nil {
        return nil, err
    }
    return &price, nil
}

func main() {
    app := iris.New()
    monitorSvc := NewMonitorService()

    // Price endpoint that accepts a product ID and returns the current price.
    app.Get("/price/{productID:string}", func(ctx iris.Context) {
        productID := ctx.Params().Get("productID")
        // Simulate a URL from which we fetch the price.
        url := fmt.Sprintf("http://example.com/price/%s", productID)
        price, err := monitorSvc.FetchPrice(url)
        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        ctx.JSON(price)
    })

    // Start the server.
    if err := app.Listen(":8080"); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}

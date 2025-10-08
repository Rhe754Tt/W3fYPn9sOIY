// 代码生成时间: 2025-10-09 03:17:20
 * It provides an endpoint to report medical quality metrics and another to retrieve the latest metrics.
 */

package main

import (
    "fmt"
    "log"
    "time"

    "github.com/kataras/iris/v12"
)

// MedicalQualityMetric represents the data structure for medical quality metrics.
type MedicalQualityMetric struct {
    Timestamp  time.Time `json:"timestamp"`
    MetricName string    `json:"metricName"`
    MetricValue int      `json:"metricValue"`
}

var metrics = make([]MedicalQualityMetric, 0)

// addMetric is an endpoint handler that adds a new medical quality metric.
func addMetric(ctx iris.Context) {
    var metric MedicalQualityMetric
    if err := ctx.ReadJSON(&metric); err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.NewMap{"error": "Invalid JSON"})
        return
    }

    // Add the metric to the list.
    metrics = append(metrics, metric)
    ctx.JSON(iris.NewMap{
        "message": "Metric added successfully",
        "metric": metric,
    })
}

// getLatestMetrics is an endpoint handler that returns the latest medical quality metrics.
func getLatestMetrics(ctx iris.Context) {
    ctx.JSON(iris.NewMap{
        "latestMetrics": metrics,
    })
}

func main() {
    app := iris.New()

    // Define routes.
    app.Post("/metrics", addMetric)
    app.Get("/metrics", getLatestMetrics)

    // Start the server.
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}

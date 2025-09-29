// 代码生成时间: 2025-09-29 20:22:01
package main

import (
    "encoding/json"
    "fmt"
    "github.com/kataras/iris/v12"
    "log"
)

// EpidemicData represents the structure of epidemic data.
type EpidemicData struct {
    ID        string `json:"id"`
    Disease   string `json:"disease"`
    Confirmed int    `json:"confirmed"`
    Recovered int    `json:"recovered"`
    Deaths    int    `json:"deaths"`
}

// EpidemicMonitor handles the epidemic monitoring functionality.
type EpidemicMonitor struct {
    Data []EpidemicData
}

// NewEpidemicMonitor creates a new EpidemicMonitor instance.
func NewEpidemicMonitor() *EpidemicMonitor {
    return &EpidemicMonitor{
        Data: []EpidemicData{
            {ID: "1", Disease: "Flu", Confirmed: 100, Recovered: 80, Deaths: 5},
            {ID: "2", Disease: "COVID-19", Confirmed: 500, Recovered: 300, Deaths: 20},
            // Add more diseases as needed
        },
    }
}

// GetAll returns all epidemic data in JSON format.
func (m *EpidemicMonitor) GetAll(ctx iris.Context) {
    response, err := json.Marshal(m.Data)
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.WriteString("Internal Server Error")
        return
    }
    ctx.JSON(iris.StatusOK, response)
}

// GetOne returns epidemic data for a specific disease.
func (m *EpidemicMonitor) GetOne(ctx iris.Context) {
    id := ctx.Params().Get("id")
    for _, data := range m.Data {
        if data.ID == id {
            response, err := json.Marshal(data)
            if err != nil {
                ctx.StatusCode(iris.StatusInternalServerError)
                ctx.WriteString("Internal Server Error")
                return
            }
            ctx.JSON(iris.StatusOK, response)
            return
        }
    }
    ctx.StatusCode(iris.StatusNotFound)
    ctx.WriteString("Epidemic data not found")
}

func main() {
    app := iris.New()

    // Create an instance of EpidemicMonitor.
    monitor := NewEpidemicMonitor()

    // Define routes for epidemic monitoring.
    app.Get("/epidemics", monitor.GetAll)
    app.Get("/epidemics/{id}", monitor.GetOne)

    // Start the Iris server.
    log.Printf("Server is running at :8080")
    if err := app.Run(iris.Addr(":8080")); err != nil {
        log.Fatalf("Server failed to start: %s", err)
    }
}
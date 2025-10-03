// 代码生成时间: 2025-10-03 20:21:57
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "net/http"
)

// IoTGateway represents an IoT gateway
type IoTGateway struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Location string `json:"location"`
}

// gateways is a mock database for storing IoT gateways
var gateways = []IoTGateway{
    {ID: "1", Name: "Gateway A", Location: "Location A"},
    {ID: "2", Name: "Gateway B", Location: "Location B"},
}

// getAllGateways handles GET requests for listing all IoT gateways
func getAllGateways(ctx iris.Context) {
    gatewaysJSON, err := json.Marshal(gateways)
    if err != nil {
        ctx.StatusCode(http.StatusInternalServerError)
        ctx.WriteString(fmt.Sprintf("Error marshaling gateways: %s", err.Error()))
        return
    }
    ctx.JSON(iris.StatusOK, string(gatewaysJSON))
}

// getGatewayByID handles GET requests for retrieving an IoT gateway by ID
func getGatewayByID(ctx iris.Context) {
    id := ctx.Params().Get("id")
    for _, gateway := range gateways {
        if gateway.ID == id {
            ctx.JSON(iris.StatusOK, gateway)
            return
        }
    }
    ctx.StatusCode(http.StatusNotFound)
    ctx.WriteString(fmt.Sprintf("Gateway with ID %s not found", id))
}

// addGateway handles POST requests for adding a new IoT gateway
func addGateway(ctx iris.Context) {
    var newGateway IoTGateway
    if err := ctx.ReadJSON(&newGateway); err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.WriteString(fmt.Sprintf("Error reading JSON: %s", err.Error()))
        return
    }
    gateways = append(gateways, newGateway)
    ctx.JSON(iris.StatusOK, newGateway)
}

// updateGateway handles PUT requests for updating an IoT gateway
func updateGateway(ctx iris.Context) {
    id := ctx.Params().Get("id")
    var updatedGateway IoTGateway
    if err := ctx.ReadJSON(&updatedGateway); err != nil {
        ctx.StatusCode(http.StatusBadRequest)
        ctx.WriteString(fmt.Sprintf("Error reading JSON: %s", err.Error()))
        return
    }
    for i, gateway := range gateways {
        if gateway.ID == id {
            gateways[i] = updatedGateway
            ctx.JSON(iris.StatusOK, updatedGateway)
            return
        }
    }
    ctx.StatusCode(http.StatusNotFound)
    ctx.WriteString(fmt.Sprintf("Gateway with ID %s not found", id))
}

// deleteGateway handles DELETE requests for deleting an IoT gateway
func deleteGateway(ctx iris.Context) {
    id := ctx.Params().Get("id")
    for i, gateway := range gateways {
        if gateway.ID == id {
            gateways = append(gateways[:i], gateways[i+1:]...)
            ctx.StatusCode(http.StatusOK)
            ctx.WriteString("Gateway deleted successfully")
            return
        }
    }
    ctx.StatusCode(http.StatusNotFound)
    ctx.WriteString(fmt.Sprintf("Gateway with ID %s not found", id))
}

func main() {
    app := iris.New()
    app.Logger().SetLevel("github.com/kataras/iris/v12".LoggerLevelSilent) // Disable logging for simplicity

    // Define routes
    app.Get("/gateways", getAllGateways)
    app.Get("/gateways/{id}", getGatewayByID)
    app.Post("/gateways", addGateway)
    app.Put("/gateways/{id}", updateGateway)
    app.Delete("/gateways/{id}", deleteGateway)

    // Start the server
    app.Listen(":8080")
}
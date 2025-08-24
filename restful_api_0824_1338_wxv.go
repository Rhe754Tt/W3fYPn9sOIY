// 代码生成时间: 2025-08-24 13:38:10
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "github.com/kataras/iris/v12/mvc"
)

// APIResponse is a struct for our API response
type APIResponse struct {
    Message string `json:"message"`
    Data    interface{} `json:"data"`
}

// NewAPIResponse creates a new APIResponse with a message
func NewAPIResponse(message string, data interface{}) APIResponse {
    return APIResponse{
        Message: message,
        Data: data,
    }
}

// MainController is our main controller
type MainController struct {
    // no fields needed for now
}

// Get is a method for the GET request
func (c *MainController) Get() mvc.Result {
    // Respond with a simple message and a data payload
    return mvc.Response{
        StatusCode: iris.StatusOK,
        ContentType: "application/json",
        Object: NewAPIResponse("Hello from Iris!", map[string]string{
            "version": "v12",
            "controller": "MainController",
        })},
}

// main is the entry point of our application
func main() {
    app := iris.New()
    
    // Register our controller
    mainCtrl := MainController{}
    app.Party("/api").Handle(&mainCtrl)
    
    // Start the server
    fmt.Println("Server is running at :8080")
    if err := app.Listen(":8080"); err != nil {
        fmt.Println("Error serving the app: ", err)
    }
}

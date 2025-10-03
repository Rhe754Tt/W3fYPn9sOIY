// 代码生成时间: 2025-10-04 03:46:24
package main

import (
    "fmt"
    "math"
    "net/http"

    "github.com/kataras/iris/v12"
)

// Collision represents a collision between two objects
type Collision struct {
    ObjectA string `json:"objectA"`
    ObjectB string `json:"objectB"`
}

// detectCollision checks if two objects collide based on their bounding boxes
func detectCollision(objA, objB BoundingBox) (bool, error) {
    // Check if either object's bounding box is not valid
    if !objA.IsValid() || !objB.IsValid() {
        return false, fmt.Errorf("invalid bounding box")
    }

    // Check for overlap in each dimension
    if objA.Min.X < objB.Max.X && objA.Max.X > objB.Min.X &&
        objA.Min.Y < objB.Max.Y && objA.Max.Y > objB.Min.Y {
        return true, nil
    }
    return false, nil
}

// BoundingBox represents the bounding box of an object
type BoundingBox struct {
    Min, Max Point
}

// Point represents a point in 2D space
type Point struct {
    X, Y float64
}

// IsValid checks if the bounding box is valid (non-zero width and height)
func (b *BoundingBox) IsValid() bool {
    return b.Min.X < b.Max.X && b.Min.Y < b.Max.Y
}

func main() {
    app := iris.New()
    app.RegisterView(iris.HTML("./templates", ".html"))

    // Define a route for collision detection
    app.Post("/detect", func(ctx iris.Context) {
        // Parse the request body as two bounding boxes
        var boundingBoxA, boundingBoxB BoundingBox
        if err := ctx.ReadJSON(&boundingBoxA); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }
        if err := ctx.ReadJSON(&boundingBoxB); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        // Detect collision
        collided, err := detectCollision(boundingBoxA, boundingBoxB)
        if err != nil {
            ctx.StatusCode(http.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
            return
        }

        // Respond with the collision result
        ctx.JSON(iris.Map{
            "collision": collided,
        })
    })

    // Start the Iris server
    if err := app.Run(iris.Addr(":8080")); err != nil {
        fmt.Println(err)
    }
}

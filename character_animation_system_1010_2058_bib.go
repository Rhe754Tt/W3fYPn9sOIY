// 代码生成时间: 2025-10-10 20:58:57
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// Animation represents a simple structure for an animation
type Animation struct {
    ID        string
    Name      string
    Loop      bool
    FrameDelay int
}

// AnimationService is responsible for handling animation-related operations
type AnimationService struct {
    Animations map[string]Animation
}

// NewAnimationService creates a new AnimationService instance
func NewAnimationService() *AnimationService {
    return &AnimationService{
        Animations: make(map[string]Animation),
    }
}

// AddAnimation adds a new animation to the service
func (s *AnimationService) AddAnimation(animation Animation) error {
    if _, exists := s.Animations[animation.ID]; exists {
        return fmt.Errorf("animation with id %s already exists", animation.ID)
    }
    s.Animations[animation.ID] = animation
    return nil
}

// GetAnimation retrieves an animation by its ID
func (s *AnimationService) GetAnimation(animationID string) (*Animation, error) {
    animation, exists := s.Animations[animationID]
    if !exists {
        return nil, fmt.Errorf("animation with id %s not found", animationID)
    }
    return &animation, nil
}

// Character represents an entity that can have animations
type Character struct {
    ID     string
    Name   string
    Animations []Animation
}

// AnimationController handles HTTP requests related to animations
type AnimationController struct {
    Service *AnimationService
}

// NewAnimationController creates a new AnimationController instance
func NewAnimationController(service *AnimationService) *AnimationController {
    return &AnimationController{
        Service: service,
    }
}

// AddAnimationHandler handles the HTTP request to add a new animation
func (c *AnimationController) AddAnimationHandler(ctx iris.Context) {
    var animation Animation
    if err := ctx.ReadJSON(&animation); err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
        ctx.JSON(iris.Map{
            "error": "Invalid request body",
        })
        return
    }
    if err := c.Service.AddAnimation(animation); err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(iris.Map{
        "message": "Animation added successfully",
    })
}

// GetAnimationHandler handles the HTTP request to get an animation
func (c *AnimationController) GetAnimationHandler(ctx iris.Context) {
    animationID := ctx.URLParam("id")
    animation, err := c.Service.GetAnimation(animationID)
    if err != nil {
        ctx.StatusCode(iris.StatusInternalServerError)
        ctx.JSON(iris.Map{
            "error": err.Error(),
        })
        return
    }
    ctx.JSON(animation)
}

func main() {
    app := iris.New()
    service := NewAnimationService()
    controller := NewAnimationController(service)

    // Define routes
    app.Post("/animations", controller.AddAnimationHandler)
    app.Get("/animations/{id}", controller.GetAnimationHandler)

    // Start the server
    app.Listen(":8080")
}

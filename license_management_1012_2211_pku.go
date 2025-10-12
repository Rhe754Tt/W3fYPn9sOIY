// 代码生成时间: 2025-10-12 22:11:57
package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/kataras/iris/v12"
)

// License represents the structure of a license
type License struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Status   string `json:"status"`
    Expires  string `json:"expires"` // ISO 8601 formatted date
    CreatedAt string `json:"createdAt"` // ISO 8601 formatted date
}

// LicenseService provides functionality to manage licenses
type LicenseService struct {
    // Storage for licenses, in a real-world scenario this would be a database
    licenses map[string]License
}

// NewLicenseService creates a new LicenseService
func NewLicenseService() *LicenseService {
    return &LicenseService{
        licenses: make(map[string]License),
    }
}

// AddLicense adds a new license to the storage
func (s *LicenseService) AddLicense(l License) error {
    if _, exists := s.licenses[l.ID]; exists {
        return fmt.Errorf("license with ID %s already exists", l.ID)
    }
    s.licenses[l.ID] = l
    return nil
}

// GetLicense retrieves a license by its ID
func (s *LicenseService) GetLicense(id string) (License, error) {
    l, exists := s.licenses[id]
    if !exists {
        return License{}, fmt.Errorf("license with ID %s not found", id)
    }
    return l, nil
}

// UpdateLicense updates an existing license
func (s *LicenseService) UpdateLicense(id string, l License) error {
    if _, exists := s.licenses[id]; !exists {
        return fmt.Errorf("license with ID %s not found", id)
    }
    s.licenses[id] = l
    return nil
}

// DeleteLicense deletes a license by its ID
func (s *LicenseService) DeleteLicense(id string) error {
    if _, exists := s.licenses[id]; !exists {
        return fmt.Errorf("license with ID %s not found", id)
    }
    delete(s.licenses, id)
    return nil
}

func main() {
    app := iris.New()
    service := NewLicenseService()

    // Licenses API
    app.Post("/licenses", func(ctx iris.Context) {
        var l License
        if err := ctx.ReadJSON(&l); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }
        if err := service.AddLicense(l); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }
        ctx.StatusCode(http.StatusCreated)
        ctx.JSON(iris.Map{
            "message": "License created successfully",
            "license": l,
        })
    })

    app.Get("/licenses/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        l, err := service.GetLicense(id)
        if err != nil {
            ctx.StatusCode(http.StatusNotFound)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }
        ctx.JSON(l)
    })

    app.Put("/licenses/{id}", func(ctx iris.Context) {
        var l License
        id := ctx.Params().Get("id")
        if err := ctx.ReadJSON(&l); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }
        l.ID = id // Ensure ID is set on update
        if err := service.UpdateLicense(id, l); err != nil {
            ctx.StatusCode(http.StatusNotFound)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }
        ctx.JSON(l)
    })

    app.Delete("/licenses/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        if err := service.DeleteLicense(id); err != nil {
            ctx.StatusCode(http.StatusNotFound)
            ctx.JSON(iris.Map{"error": err.Error()})
            return
        }
        ctx.JSON(iris.Map{
            "message": "License deleted successfully",
        })
    })

    // Start the Iris server
    log.Printf("Server is running on http://localhost:8080")
    if err := app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8")); err != nil {
        log.Fatalf("An error occurred while running the server: %v", err)
    }
}
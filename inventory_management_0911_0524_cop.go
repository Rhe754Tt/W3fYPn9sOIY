// 代码生成时间: 2025-09-11 05:24:09
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
    "log"
)

// InventoryItem represents an item in the inventory.
type InventoryItem struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Quantity  int    `json:"quantity"`
}

// InventoryService is the service managing the inventory operations.
type InventoryService struct {
    items map[string]InventoryItem
}

// NewInventoryService creates a new InventoryService.
func NewInventoryService() *InventoryService {
    return &InventoryService{
        items: make(map[string]InventoryItem),
    }
}

// AddItem adds a new item to the inventory or updates an existing one.
func (s *InventoryService) AddItem(item InventoryItem) error {
    s.items[item.ID] = item
    return nil
}

// GetItem retrieves an item by its ID.
func (s *InventoryService) GetItem(id string) (*InventoryItem, error) {
    item, exists := s.items[id]
    if !exists {
        return nil, fmt.Errorf("item with id %s not found", id)
    }
    return &item, nil
}

// UpdateQuantity updates the quantity of an item in the inventory.
func (s *InventoryService) UpdateQuantity(id string, quantity int) error {
    if item, exists := s.items[id]; exists {
        item.Quantity = quantity
        s.items[id] = item
        return nil
    }
    return fmt.Errorf("item with id %s not found", id)
}

// DeleteItem removes an item from the inventory.
func (s *InventoryService) DeleteItem(id string) error {
    if _, exists := s.items[id]; exists {
        delete(s.items, id)
        return nil
    }
    return fmt.Errorf("item with id %s not found", id)
}

func main() {
    app := iris.New()
    inventoryService := NewInventoryService()

    // Add sample items to the inventory for demonstration.
    err := inventoryService.AddItem(InventoryItem{ID: "1", Name: "Apple", Quantity: 10})
    if err != nil {
        log.Fatalf("Error adding item: %v", err)
    }
    err = inventoryService.AddItem(InventoryItem{ID: "2", Name: "Banana", Quantity: 20})
    if err != nil {
        log.Fatalf("Error adding item: %v", err)
    }

    // API endpoint to get an item by ID.
    app.Get("/item/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        item, err := inventoryService.GetItem(id)
        if err != nil {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
        } else {
            ctx.JSON(item)
        }
    })

    // API endpoint to update the quantity of an item.
    app.Post("/item/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        var newItem InventoryItem
        if err := ctx.ReadJSON(&newItem); err != nil {
            ctx.StatusCode(iris.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "invalid request body",
            })
            return
        }
        err := inventoryService.UpdateQuantity(id, newItem.Quantity)
        if err != nil {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
        } else {
            ctx.JSON(iris.Map{
                "message": "item updated successfully",
            })
        }
    })

    // API endpoint to delete an item.
    app.Delete("/item/{id}", func(ctx iris.Context) {
        id := ctx.Params().Get("id")
        err := inventoryService.DeleteItem(id)
        if err != nil {
            ctx.StatusCode(iris.StatusNotFound)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
        } else {
            ctx.JSON(iris.Map{
                "message": "item deleted successfully",
            })
        }
    })

    // Start the Iris HTTP server.
    if err := app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8")); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
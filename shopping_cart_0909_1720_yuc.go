// 代码生成时间: 2025-09-09 17:20:05
package main

import (
    "fmt"
    "github.com/kataras/iris/v12"
)

// ShoppingCart 表示购物车结构
type ShoppingCart struct {
    ID        string         `json:"id"`
    Items     []*CartItem    `json:"items"`
    TotalCost float64        `json:"total_cost"`
}

// CartItem 表示购物车中的商品项
type CartItem struct {
    ProductID string  `json:"product_id"`
    Quantity  int     `json:"quantity"`
    Price     float64 `json:"price"`
}

// NewShoppingCart 创建一个新的购物车
func NewShoppingCart(id string) *ShoppingCart {
    return &ShoppingCart{
        ID: id,
        Items: make([]*CartItem, 0),
    }
}

// AddItem 向购物车添加商品项
func (cart *ShoppingCart) AddItem(productID string, quantity int, price float64) error {
    if quantity <= 0 || price <= 0 {
        return fmt.Errorf("invalid quantity or price")
    }

    var item *CartItem
    for _, item = range cart.Items {
        if item.ProductID == productID {
            item.Quantity += quantity
            return nil
        }
    }

    item = &CartItem{
        ProductID: productID,
        Quantity:  quantity,
        Price:    price,
    }
    cart.Items = append(cart.Items, item)
    return nil
}

// CalculateTotalCost 计算购物车的总成本
func (cart *ShoppingCart) CalculateTotalCost() error {
    cart.TotalCost = 0
    for _, item := range cart.Items {
        cart.TotalCost += item.Price * float64(item.Quantity)
    }
    return nil
}

func main() {
    app := iris.New()

    // 购物车实例
    cart := NewShoppingCart("cart1")

    // 添加商品路由
    app.Post("/add", func(ctx iris.Context) {
        productID := ctx.URLParam("product_id")
        quantityStr := ctx.URLParam("quantity")
        quantity, _ := strconv.Atoi(quantityStr)
        priceStr := ctx.URLParam("price")
        price, _ := strconv.ParseFloat(priceStr, 64)

        err := cart.AddItem(productID, quantity, price)

        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
        } else {
            ctx.JSON(iris.Map{
                "message": "Item added successfully",
            })
        }
    })

    // 计算总成本路由
    app.Get("/total", func(ctx iris.Context) {
        err := cart.CalculateTotalCost()

        if err != nil {
            ctx.StatusCode(iris.StatusInternalServerError)
            ctx.JSON(iris.Map{
                "error": err.Error(),
            })
        } else {
            ctx.JSON(iris.Map{
                "total_cost": cart.TotalCost,
            })
        }
    })

    // 启动服务器
    app.Listen(":8080")
}
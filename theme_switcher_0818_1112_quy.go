// 代码生成时间: 2025-08-18 11:12:04
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/kataras/iris/v12"
)

// ThemeManager is a struct that holds the current theme.
type ThemeManager struct {
    currentTheme string
}

// NewThemeManager creates a new ThemeManager with a default theme.
func NewThemeManager(defaultTheme string) *ThemeManager {
    return &ThemeManager{
        currentTheme: defaultTheme,
    }
}

// SwitchTheme changes the current theme to the given theme.
func (tm *ThemeManager) SwitchTheme(theme string) error {
    if theme != "light" && theme != "dark" {
        return fmt.Errorf("invalid theme: %s", theme)
    }
    tm.currentTheme = theme
    return nil
}

// GetCurrentTheme returns the current theme.
func (tm *ThemeManager) GetCurrentTheme() string {
    return tm.currentTheme
}

func main() {
    // Initialize the theme manager with a default theme.
    tm := NewThemeManager("light")

    app := iris.New()

    // Set up a route for switching the theme.
    app.Post("/switch-theme", func(ctx iris.Context) {
        var themeData struct {
            Theme string `json:"theme"`
        }
        if err := ctx.ReadJSON(&themeData); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": fmt.Sprintf("Invalid request: %s", err),
            })
            return
        }

        if err := tm.SwitchTheme(themeData.Theme); err != nil {
            ctx.StatusCode(http.StatusBadRequest)
            ctx.JSON(iris.Map{
                "error": "Invalid theme provided",
            })
            return
        }

        ctx.StatusCode(http.StatusOK)
        ctx.JSON(iris.Map{
            "message": "Theme switched successfully",
            "currentTheme": tm.GetCurrentTheme(),
        })
    })

    // Set up a route to get the current theme.
    app.Get("/current-theme", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "currentTheme": tm.GetCurrentTheme(),
        })
    })

    // Start the Iris web server.
    if err := app.Run(iris.Addr:":8080"); err != nil {
        log.Fatalf("Failed to run the server: %s", err)
    }
}
// 代码生成时间: 2025-10-07 02:11:19
package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "fmt"
)

// User represents a user with ID, Name, Age, and Email
type User struct {
    gorm.Model
    Name  string
    Age   int
    Email string `gorm:"type:varchar(100);uniqueIndex"`
}

func main() {
    // Initialize a new SQLite database connection
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&User{})

    // Create a new user
    user := User{Name: "John Doe", Age: 28, Email: "johndoe@example.com"}
    if err := db.Create(&user).Error; err != nil {
        panic("failed to create user")
    }

    // Read a user with a specific ID
    var user1 User
    db.First(&user1, 1) // find with id 1
    fmt.Println(user1)

    // Update user's name and age
    db.Model(&user1).Update("Name", "Jane Doe")
    db.Model(&user1).Update("Age", 30)

    // Delete a user
    db.Delete(&user1, 1)
}

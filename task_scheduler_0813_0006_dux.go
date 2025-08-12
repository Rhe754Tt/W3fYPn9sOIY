// 代码生成时间: 2025-08-13 00:06:16
// task_scheduler.go
package main

import (
    "fmt"
    "time"
    "github.com/kataras/iris/v12"
)

// Scheduler represents a task scheduler
type Scheduler struct {
    tasks map[string]*Task
}

// Task defines the structure of a task
type Task struct {
    ID     string
    Action func() error
}

// NewScheduler creates a new Scheduler instance
func NewScheduler() *Scheduler {
    return &Scheduler{
        tasks: make(map[string]*Task),
    }
}

// AddTask adds a new task to the scheduler
func (s *Scheduler) AddTask(id string, f func() error) {
    s.tasks[id] = &Task{
        ID:   id,
        Action: f,
    }
}

// Run starts the scheduler, executing tasks based on their schedules
func (s *Scheduler) Run() {
    for id, task := range s.tasks {
        go func(id string, task *Task) {
            if err := task.Action(); err != nil {
                fmt.Printf("Error executing task %s: %s
", id, err)
            }
        }(id, task)
    }
}

func main() {
    // Create a new Iris web server
    app := iris.New()
    scheduler := NewScheduler()

    // Define a sample task
    err := scheduler.AddTask("sampleTask", func() error {
        // Simulate a task that might fail
        time.Sleep(2 * time.Second)
        return fmt.Errorf("sample task failed")
    })
    if err != nil {
        fmt.Println("Error adding task: ", err)
        return
    }

    // Start the scheduler
    scheduler.Run()

    // Start the Iris server on port 8080
    app.Listen(":8080")
}

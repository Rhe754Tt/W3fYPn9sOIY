// 代码生成时间: 2025-08-12 19:39:49
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "time"
)

// LogEntry represents a structured log entry with its timestamp and message.
type LogEntry struct {
    Timestamp time.Time
    Message   string
}

// parseLogLine takes a line from a log file and parses it into a LogEntry struct.
// It assumes the log format is something like: [2006-01-02 15:04:05] Some message here.
func parseLogLine(line string) (*LogEntry, error) {
    // Split the line by the first ']', which should separate the timestamp and the message.
    parts := strings.SplitN(line, "]", 2)
    if len(parts) < 2 {
        return nil, fmt.Errorf("invalid log line format: %s", line)
    }

    // Extract the timestamp and clean it from brackets.
    timestampStr := strings.TrimSpace(parts[0][1:]) // Remove the '['
    // Extract the message.
    message := strings.TrimSpace(parts[1])

    // Parse the timestamp.
    timestamp, err := time.Parse("2006-01-02 15:04:05", timestampStr)
    if err != nil {
        return nil, err
    }

    return &LogEntry{Timestamp: timestamp, Message: message}, nil
}

// parseLogFile reads a log file and parses all its entries, returning a slice of LogEntry structs.
func parseLogFile(filePath string) ([]LogEntry, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var entries []LogEntry

    for scanner.Scan() {
        line := scanner.Text()
        entry, err := parseLogLine(line)
        if err != nil {
            fmt.Printf("Skipping invalid log line: %s
", err)
            continue
        }
        entries = append(entries, *entry)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return entries, nil
}

// main function to demonstrate the log file parsing.
func main() {
    // Replace with the actual path to your log file.
    logFilePath := "./logfile.log"
    entries, err := parseLogFile(logFilePath)
    if err != nil {
        fmt.Printf("Error parsing log file: %s
", err)
        return
    }

    // Print out the parsed log entries.
    for _, entry := range entries {
        fmt.Printf("Timestamp: %s, Message: %s
", entry.Timestamp.Format("2006-01-02 15:04:05"), entry.Message)
    }
}
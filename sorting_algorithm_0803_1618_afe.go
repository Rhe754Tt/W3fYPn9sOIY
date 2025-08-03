// 代码生成时间: 2025-08-03 16:18:57
package main

import (
    "fmt"
    "sort"
)

// BubbleSort performs bubble sort algorithm on a slice of integers.
// It takes a slice of integers and returns a sorted slice.
func BubbleSort(numbers []int) ([]int, error) {
    if numbers == nil {
        return nil, fmt.Errorf("input slice cannot be nil")
    }
    for i := 0; i < len(numbers); i++ {
        for j := 0; j < len(numbers)-i-1; j++ {
            if numbers[j] > numbers[j+1] {
                // Swap the elements
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
            }
        }
    }
    return numbers, nil
}

// SelectionSort performs selection sort algorithm on a slice of integers.
// It takes a slice of integers and returns a sorted slice.
func SelectionSort(numbers []int) ([]int, error) {
    if numbers == nil {
        return nil, fmt.Errorf("input slice cannot be nil")
    }
    for i := 0; i < len(numbers); i++ {
        minIndex := i
        for j := i + 1; j < len(numbers); j++ {
            if numbers[j] < numbers[minIndex] {
                minIndex = j
            }
        }
        if minIndex != i {
            numbers[i], numbers[minIndex] = numbers[minIndex], numbers[i]
        }
    }
    return numbers, nil
}

// InsertionSort performs insertion sort algorithm on a slice of integers.
// It takes a slice of integers and returns a sorted slice.
func InsertionSort(numbers []int) ([]int, error) {
    if numbers == nil {
        return nil, fmt.Errorf("input slice cannot be nil")
    }
    for i := 1; i < len(numbers); i++ {
        key := numbers[i]
        j := i - 1
        for j >= 0 && numbers[j] > key {
            numbers[j+1] = numbers[j]
            j = j - 1
        }
        numbers[j+1] = key
    }
    return numbers, nil
}

func main() {
    numbers := []int{64, 34, 25, 12, 22, 11, 90}

    // Example of using BubbleSort
    sortedNumbers, err := BubbleSort(numbers)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sorted Numbers (Bubble Sort): ", sortedNumbers)
    }

    // Example of using SelectionSort
    sortedNumbers, err = SelectionSort(numbers)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sorted Numbers (Selection Sort): ", sortedNumbers)
    }

    // Example of using InsertionSort
    sortedNumbers, err = InsertionSort(numbers)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Sorted Numbers (Insertion Sort): ", sortedNumbers)
    }
}

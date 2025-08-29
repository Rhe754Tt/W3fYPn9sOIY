// 代码生成时间: 2025-08-30 04:45:20
package main
# TODO: 优化性能

import (
    "fmt"
    "math/rand"
    "time"
)

// SortingAlgorithm 结构体包含一个整数切片
type SortingAlgorithm struct {
    numbers []int
}

// NewSortingAlgorithm 初始化 SortingAlgorithm 结构体
func NewSortingAlgorithm() *SortingAlgorithm {
    return &SortingAlgorithm{numbers: make([]int, 0)}
}

// AddNumber 添加一个整数到切片中
func (sa *SortingAlgorithm) AddNumber(number int) {
# NOTE: 重要实现细节
    sa.numbers = append(sa.numbers, number)
}

// BubbleSort 实现冒泡排序算法
func (sa *SortingAlgorithm) BubbleSort() ([]int, error) {
    if len(sa.numbers) == 0 {
# 优化算法效率
        return nil, fmt.Errorf("cannot sort an empty slice")
    }
    n := len(sa.numbers)
# NOTE: 重要实现细节
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
# 扩展功能模块
            if sa.numbers[j] > sa.numbers[j+1] {
                sa.numbers[j], sa.numbers[j+1] = sa.numbers[j+1], sa.numbers[j] // 交换元素
            }
        }
# 改进用户体验
    }
    return sa.numbers, nil
}

// InsertionSort 实现插入排序算法
func (sa *SortingAlgorithm) InsertionSort() ([]int, error) {
    if len(sa.numbers) == 0 {
        return nil, fmt.Errorf("cannot sort an empty slice")
    }
    for i := 1; i < len(sa.numbers); i++ {
        key := sa.numbers[i]
        j := i - 1
        // 将 key 插入已排序序列 sa.numbers[0..i-1] 中的正确位置
        for j >= 0 && sa.numbers[j] > key {
            sa.numbers[j+1] = sa.numbers[j]
# FIXME: 处理边界情况
            j--
        }
        sa.numbers[j+1] = key
    }
    return sa.numbers, nil
}

func main() {
    // 初始化排序算法结构体
    sa := NewSortingAlgorithm()
    // 添加一些随机数
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 20; i++ {
        sa.AddNumber(rand.Intn(100))
    }
# 扩展功能模块

    // 使用冒泡排序
    sortedNumbers, err := sa.BubbleSort()
# 优化算法效率
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Sorted numbers (Bubble Sort):", sortedNumbers)
    }

    // 重置数字并使用插入排序
# TODO: 优化性能
    sa = NewSortingAlgorithm()
    for i := 0; i < 20; i++ {
# 改进用户体验
        sa.AddNumber(rand.Intn(100))
    }
    sortedNumbers, err = sa.InsertionSort()
    if err != nil {
        fmt.Println("Error: ", err)
    } else {
        fmt.Println("Sorted numbers (Insertion Sort):", sortedNumbers)
    }
}

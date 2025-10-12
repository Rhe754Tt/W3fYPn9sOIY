// 代码生成时间: 2025-10-13 02:38:33
package main

import (
    "fmt"
    "math"

    "github.com/iris-contrib/middleware/cors"
    "github.com/kataras/iris/v12"
)

// ClusterAnalysis 结构体用于存储聚类分析相关的数据和方法
type ClusterAnalysis struct {
    // 数据点的集合
    points []Point
}

// Point 结构体表示单个数据点
type Point struct {
    X float64
    Y float64
}

// NewClusterAnalysis 初始化聚类分析工具
func NewClusterAnalysis(points []Point) *ClusterAnalysis {
    return &ClusterAnalysis{
        points: points,
    }
}

// KMeans 执行K均值聚类算法
func (ca *ClusterAnalysis) KMeans(k int) ([]int, error) {
    if k <= 0 || k > len(ca.points) {
        return nil, fmt.Errorf("k value is out of range")
    }

    // 初始化聚类中心
    centroids := make([]Point, k)
    for i := range centroids {
        centroids[i] = ca.points[i]
    }

    clusters := make([][]Point, k)
    for i := range clusters {
        clusters[i] = []Point{centroids[i]}
    }

    for !isStable(centroids, clusters) {
        assignPointsToClusters(ca.points, centroids, clusters)
        newCentroids := calculateNewCentroids(clusters)
        if len(newCentroids) != len(centroids) {
            return nil, fmt.Errorf("could not calculate new centroids")
        }
        centroids = newCentroids
    }

    clusterAssignments := make([]int, len(ca.points))
    for i, cluster := range clusters {
        for j := range cluster {
            clusterAssignments[j] = i
        }
    }

    return clusterAssignments, nil
}

// isStable 检查聚类是否稳定
func isStable(centroids []Point, clusters [][]Point) bool {
    for i, cluster := range clusters {
        newCentroid := meanPoint(cluster)
        if centroids[i].X != newCentroid.X || centroids[i].Y != newCentroid.Y {
            return false
        }
    }
    return true
}

// assignPointsToClusters 将点分配到最近的聚类中心
func assignPointsToClusters(points []Point, centroids []Point, clusters [][]Point) {
    for _, point := range points {
        nearestCentroidIndex := findNearestCentroid(point, centroids)
        clusters[nearestCentroidIndex] = append(clusters[nearestCentroidIndex], point)
    }
}

// findNearestCentroid 查找最近的聚类中心
func findNearestCentroid(point Point, centroids []Point) int {
    var nearestIndex int
    var nearestDistance float64 = math.MaxFloat64
    for i, centroid := range centroids {
        distance := distanceBetweenPoints(point, centroid)
        if distance < nearestDistance {
            nearestDistance = distance
            nearestIndex = i
        }
    }
    return nearestIndex
}

// distanceBetweenPoints 计算两点之间的距离
func distanceBetweenPoints(p1, p2 Point) float64 {
    return math.Sqrt(math.Pow(p1.X-p2.X, 2) + math.Pow(p1.Y-p2.Y, 2))
}

// meanPoint 计算聚类中点的平均值
func meanPoint(points []Point) Point {
    sumX := 0.0
    sumY := 0.0
    for _, point := range points {
        sumX += point.X
        sumY += point.Y
    }
    return Point{
        X: sumX / float64(len(points)),
        Y: sumY / float64(len(points)),
    }
}

// calculateNewCentroids 计算新的聚类中心
func calculateNewCentroids(clusters [][]Point) []Point {
    newCentroids := make([]Point, len(clusters))
    for i, cluster := range clusters {
        newCentroids[i] = meanPoint(cluster)
    }
    return newCentroids
}

func main() {
    app := iris.New()
    app.Use(cors.New(cors.Options{
       AllowedOrigins: []string{iris.String("*")},
       AllowedMethods: []string{iris.MethodGet, iris.MethodPost, iris.MethodPut, iris.MethodPatch, iris.MethodDelete, iris.MethodOptions},
       AllowCredentials: true,
    }))

    // 测试数据点
    points := []Point{
        {X: 1.0, Y: 2.0},
        {X: 2.0, Y: 3.0},
        {X: 3.0, Y: 4.0},
        {X: 4.0, Y: 5.0},
    }

    // 创建聚类分析工具
    ca := NewClusterAnalysis(points)

    // 执行K均值聚类
    clusterAssignments, err := ca.KMeans(2)
    if err != nil {
        fmt.Printf("Error performing KMeans clustering: %v
", err)
        return
    }

    // 打印聚类结果
    fmt.Printf("Cluster assignments: %v
", clusterAssignments)

    app.Listen(":8080")
}

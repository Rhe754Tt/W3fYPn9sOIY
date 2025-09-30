// 代码生成时间: 2025-09-30 21:14:16
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Individual 表示遗传算法中的个体
type Individual struct {
    Genome []int // 基因组，这里使用整数数组作为示例
    Fit    float64 // 适应度
}

// Population 表示遗传算法中的种群
type Population struct {
    Pool []Individual // 种群中的个体
}

// NewPopulation 创建一个新的种群
func NewPopulation(size int) Population {
    var pool []Individual
    for i := 0; i < size; i++ {
        genome := make([]int, 10) // 假设基因组长度为10
        for j := range genome {
            genome[j] = rand.Intn(2) // 随机生成0或1
        }
        pool = append(pool, Individual{Genome: genome})
    }
    return Population{Pool: pool}
}

// Evaluate 评估种群中每个个体的适应度
func (p *Population) Evaluate() {
    for i := range p.Pool {
        p.Pool[i].Fit = evaluateIndividual(&p.Pool[i])
    }
}

// EvaluateIndividual 评估单个个体的适应度
// 这里使用一个简单的适应度函数作为示例
func evaluateIndividual(individual *Individual) float64 {
    // 简单的适应度函数，仅作为示例
    fit := float64(0)
    for _, gene := range individual.Genome {
        fit += float64(gene)
    }
    return fit
}

// Select 选择适应度最高的个体
func (p *Population) Select() *Individual {
    maxFit := float64(-1)
    bestIndividual := &p.Pool[0]
    for _, individual := range p.Pool {
        if individual.Fit > maxFit {
            maxFit = individual.Fit
            bestIndividual = &individual
        }
    }
    return bestIndividual
}

// Crossover 交叉操作，生成新的后代
func (p *Population) Crossover(parent1, parent2 *Individual) *Individual {
    child := Individual{Genome: make([]int, len(parent1.Genome))}
    for i := range child.Genome {
        child.Genome[i] = parent1.Genome[i]
        if rand.Float64() > 0.5 {
            child.Genome[i] = parent2.Genome[i]
        }
    }
    return &child
}

// Mutate 变异操作，随机改变个体的基因
func mutateIndividual(individual *Individual) {
    for i := range individual.Genome {
        if rand.Float64() < 0.1 { // 假设变异概率为10%
            individual.Genome[i] = 1 - individual.Genome[i]
        }
    }
}

// Evolve 进化操作，生成新一代种群
func (p *Population) Evolve() {
    newPool := []Individual{*p.Select()} // 保留最佳个体
    for i := 1; i < len(p.Pool); i++ {
        parent1 := p.Select()
        parent2 := p.Select()
        child := *p.Crossover(parent1, parent2)
        mutateIndividual(&child)
        newPool = append(newPool, child)
    }
    p.Pool = newPool
}

func main() {
    rand.Seed(time.Now().UnixNano()) // 初始化随机种子
    population := NewPopulation(100) // 创建一个包含100个个体的种群
    for i := 0; i < 100; i++ { // 进行100代进化
        population.Evaluate() // 评估适应度
        population.Evolve()  // 进化操作
        fmt.Printf("Generation %d: Best Fit: %.2f
", i, population.Select().Fit)
    }
}

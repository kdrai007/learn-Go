package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	fmt.Println(costs)
	lastDay := costs[0].day
	for _, c := range costs {
		if c.day > lastDay {
			lastDay = c.day
		}
	}
	costByday := make([]float64, lastDay+1)
	for i := 0; i <= lastDay; i++ {
		for _, cost := range costs {
			if cost.day == i {
				costByday[i] += cost.value
			}
		}
	}
	fmt.Printf("size=%d  and slice=%v\n", lastDay, costByday)
	return costByday
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	// fill matrix with values
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func exerciseTwo() {
	fmt.Println("--------exercise Two----------")
	costs := []cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.1},
		{5, 2.5},
		{2, 3.6},
		{1, 2.7},
		{1, 3.3},
	}
	getCostByDay(costs)

	fmt.Println("--------exercise Two.one----------")

	matrix := createMatrix(10, 7)
	//print the matrix
	for i := 0; i < 10; i++ {
		fmt.Println(matrix[i])
	}

}

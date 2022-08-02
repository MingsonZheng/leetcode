package main

import (
	"fmt"
	"math"
)

// 剑指 Offer 47. 礼物的最大价值

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(maxValue(grid))
}

func maxValue(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m)
	}
	sum := 0
	for j := 0; j < m; j++ {
		sum += grid[0][j]
		dp[0][j] = sum
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum += grid[i][0]
		dp[i][0] = sum
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}

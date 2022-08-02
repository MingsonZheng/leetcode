package main

import (
	"fmt"
	"math"
)

// 64. 最小路径和

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	len := 0
	for i := 0; i < m; i++ {
		len += grid[i][0]
		dp[i][0] = len
	}
	len = 0
	for j := 0; j < n; j++ {
		len += grid[0][j]
		dp[0][j] = len
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

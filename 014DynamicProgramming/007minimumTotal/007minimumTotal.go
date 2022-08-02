package main

import (
	"fmt"
	"math"
)

// 120. 三角形最小路径和

func main() {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1]))) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}
	res := math.MaxInt32
	for j := 0; j < n; j++ {
		if dp[n-1][j] < res {
			res = dp[n-1][j]
		}
	}
	return res
}

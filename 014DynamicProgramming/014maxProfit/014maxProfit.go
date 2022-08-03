package main

import (
	"fmt"
	"math"
)

// 309. 最佳买卖股票时机含冷冻期

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	dp[0][3] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max3(dp[i-1][0], dp[i-1][2]-prices[i], dp[i-1][3]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = dp[i-1][1]
		dp[i][3] = int(math.Max(float64(dp[i-1][2]), float64(dp[i-1][3])))
	}
	return max4(dp[n-1][0], dp[n-1][1], dp[n-1][2], dp[n-1][3])
}

func max3(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}

func max4(a, b, c, d int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if d > max {
		max = d
	}
	return max
}

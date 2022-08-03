package main

import (
	"fmt"
	"math"
)

// 714. 买卖股票的最佳时机含手续费

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fmt.Println(maxProfit(prices, 2))
}

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0] // 第 i 天持有股票
	dp[0][1] = 0          // 第 i 天不持有股票
	for i := 1; i < n; i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]-prices[i])))
		dp[i][1] = int(math.Max(float64(dp[i-1][0]+prices[i]-fee), float64(dp[i-1][1])))
	}
	return int(math.Max(float64(dp[n-1][0]), float64(dp[n-1][1])))
}

package main

import (
	"fmt"
	"math"
)

// 322. 零钱兑换

func main() {
	coins := []int{1, 2, 5}
	fmt.Println(coinChange(coins, 11))
}

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, amount+1)
		for j := 0; j <= amount; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for c := 0; c <= amount/coins[0]; c++ {
		dp[0][c*coins[0]] = c
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= amount; j++ {
			k := j / coins[i]
			for c := 0; c <= k; c++ {
				if dp[i-1][j-c*coins[i]] != math.MaxInt32 &&
					dp[i-1][j-c*coins[i]]+c < dp[i][j] {
					dp[i][j] = dp[i-1][j-c*coins[i]] + c
				}
			}
		}
	}
	if dp[n-1][amount] == math.MaxInt32 {
		return -1
	}
	return dp[n-1][amount]
}

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
	k := len(coins)
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < k; j++ {
			if i-coins[j] >= 0 &&
				dp[i-coins[j]] != math.MaxInt32 &&
				dp[i-coins[j]]+1 < dp[i] {
				dp[i] = dp[i-coins[j]] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

package main

import "fmt"

// 518. 零钱兑换 II

func main() {
	coins := []int{1, 2, 5}
	fmt.Println(change(5, coins))
}

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}
	for c := 0; c <= amount/coins[0]; c++ {
		dp[0][c*coins[0]] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= amount; j++ {
			k := j / coins[i]
			for c := 0; c <= k; c++ {
				dp[i][j] += dp[i-1][j-c*coins[i]]
			}
		}
	}
	return dp[n-1][amount]
}

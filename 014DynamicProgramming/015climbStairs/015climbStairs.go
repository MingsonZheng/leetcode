package main

import "fmt"

// 70. 爬楼梯

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs2(2))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if i-1 >= 0 {
			dp[i] += dp[i-1]
		}
		if i-2 >= 0 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

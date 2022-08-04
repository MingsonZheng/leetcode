package main

import "fmt"

// 剑指 Offer 14- I. 剪绳子

func main() {
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(10))
}

func cuttingRope(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if dp[i] < j*dp[i-j] {
				dp[i] = j * dp[i-j]
			}
		}
	}
	return dp[n]
}

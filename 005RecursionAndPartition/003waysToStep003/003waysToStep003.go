package main

import "fmt"

// 面试题 08.01. 三步问题
// 解法 3：非递归实现

func main() {
	fmt.Println(waysToStep(3))
	fmt.Println(waysToStep(5))
	fmt.Println(waysToStep(61))
}

const mod = 1000000007

var memo = make([]int, 1000001)

func waysToStep(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 4
	for i := 4; i <= n; i++ {
		dp[i] = ((dp[i-1]+dp[i-2])%mod + dp[i-3]) % mod
	}
	return dp[n]
}

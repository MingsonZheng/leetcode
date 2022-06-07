package main

import "fmt"

// 面试题 08.01. 三步问题
// 解法 2：递归 + 备忘录

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
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = ((waysToStep(n-1) + waysToStep(n-2)%mod) + waysToStep(n-3)) % mod
	return memo[n]
}

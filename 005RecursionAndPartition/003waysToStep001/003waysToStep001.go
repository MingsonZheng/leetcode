package main

import "fmt"

// 面试题 08.01. 三步问题
// 解法 1：递归，超时

func main() {
	fmt.Println(waysToStep(3))
	fmt.Println(waysToStep(5))
	fmt.Println(waysToStep(61))
}

const mod = 1000000007

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
	return ((waysToStep(n-1) + waysToStep(n-2)%mod) + waysToStep(n-3)) % mod
}

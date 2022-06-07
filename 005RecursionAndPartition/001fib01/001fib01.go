package main

import "fmt"

// 剑指 Offer 10- I. 斐波那契数列
// 解法 1：递归 超出时间限制

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(43))
}

const mod = 1000000007

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return (fib(n-1) + fib(n-2)) % mod
}

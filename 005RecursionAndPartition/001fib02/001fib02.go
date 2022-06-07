package main

import "fmt"

// 剑指 Offer 10- I. 斐波那契数列
// 解法 2：递归 + 备忘录

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(5))
	fmt.Println(fib(43))
}

const mod = 1000000007

var memo []int

func fib(n int) int {
	memo = make([]int, n+1)
	return fib_r(n)
}

func fib_r(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = (fib_r(n-1) + fib_r(n-2)) % mod
	return memo[n]
}

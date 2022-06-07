package main

import "fmt"

// 剑指 Offer 10- II. 青蛙跳台阶问题

func main() {
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
}

const mod = 1000000007

var memo = make(map[int]int, 0)

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if v, ok := memo[n]; ok {
		return v
	}
	ret := (numWays(n-1) + numWays(n-2)) % mod
	memo[n] = ret
	return ret
}

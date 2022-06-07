package main

import "fmt"

// 面试题 08.01. 三步问题
// 解法 4：非递归实现优化

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
	a := 1
	b := 2
	c := 4
	d := 0
	for i := 4; i <= n; i++ {
		d = ((c+b)%mod + a) % mod
		a, b, c = b, c, d
	}
	return d
}

package main

import "fmt"

func main() {
	res := divingBoard(1, 2, 3)
	for _, v := range res {
		fmt.Printf("%v，", v)
	}
}

// 面试题 16.11. 跳水板
func divingBoard(shorter int, longer int, k int) []int {
	// 特殊情况处理
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{k * shorter}
	}

	result := make([]int, k+1)
	// 长板子个数：0，1，2，...，k
	for i := 0; i <= k; i++ {
		result[i] = i*longer + (k-i)*shorter
	}
	return result
}

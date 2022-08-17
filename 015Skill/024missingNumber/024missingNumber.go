package main

import "fmt"

// 面试题 17.04. 消失的数字

func main() {
	nums := []int{3, 0, 1}
	fmt.Println(missingNumber(nums))
}

func missingNumber(nums []int) int {
	n := len(nums)
	ret := 0
	for i := 0; i <= n; i++ {
		ret ^= i
	}
	for i := 0; i < n; i++ {
		ret ^= nums[i]
	}
	return ret
}

package main

import "fmt"

// 剑指 Offer 56 - I. 数组中数字出现的次数

func main() {
	nums := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(nums))
}

func singleNumbers(nums []int) []int {
	xorResult := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		xorResult ^= nums[i]
	}
	tag := 1
	for xorResult&tag == 0 {
		tag = tag << 1
	}
	a := 0
	b := 0
	for i := 0; i < n; i++ {
		if nums[i]&tag == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}

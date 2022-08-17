package main

import "fmt"

// 剑指 Offer 56 - II. 数组中数字出现的次数 II

func main() {
	nums := []int{3, 4, 3, 3}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	n := len(nums)
	bits := make([]int, 32)
	mask := 1
	for i := 0; i < 32; i++ {
		for j := 0; j < n; j++ {
			if (nums[j] & mask) != 0 {
				bits[i] = (bits[i] + 1) % 3
			}
		}
		mask <<= 1
	}
	result := 0
	mask = 1
	for i := 0; i < 32; i++ {
		if bits[i] == 1 {
			result += mask
		}
		mask <<= 1
	}
	return result
}

package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 2, 5}
	fmt.Println(isStraight(nums))
}

// 剑指 Offer 61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	dup := make([]bool, 14) // 去重用的
	min := 100
	max := -1
	for i := 0; i < 5; i++ {
		if nums[i] != 0 {
			if dup[nums[i]] {
				return false
			} else {
				dup[nums[i]] = true
			}

			if nums[i] < min {
				min = nums[i]
			}

			if nums[i] > max {
				max = nums[i]
			}
		}
	}
	return (max - min) < 5
}

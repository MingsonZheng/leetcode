package main

import "fmt"

// 283. 移动零

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	p := -1
	q := 0
	for q < len(nums) {
		if nums[q] == 0 {
			q++
			continue
		}
		if nums[q] != 0 {
			nums[p+1], nums[q] = nums[q], nums[p+1]
			p++
			q++
		}
	}
}

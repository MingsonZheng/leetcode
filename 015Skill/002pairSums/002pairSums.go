package main

import (
	"fmt"
	"sort"
)

// 面试题 16.24. 数对和

func main() {
	nums := []int{5, 6, 5}
	target := 11
	fmt.Println(pairSums(nums, target))
}

func pairSums(nums []int, target int) [][]int {
	results := make([][]int, 0)
	if len(nums) == 0 {
		return results
	}
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]+nums[j] == target {
			result := make([]int, 0)
			result = append(result, nums[i])
			result = append(result, nums[j])
			results = append(results, result)
			i++
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return results
}

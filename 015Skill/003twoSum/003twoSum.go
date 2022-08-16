package main

import (
	"fmt"
	"sort"
)

// 1. 两数之和

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	sortedNums := make([]int, n)
	for i := 0; i < n; i++ {
		sortedNums[i] = nums[i]
	}
	sort.Ints(sortedNums)
	used := make([]bool, n)
	p := 0
	q := n - 1
	for p < q {
		sum := sortedNums[p] + sortedNums[q]
		if sum == target {
			oldP := find(nums, used, sortedNums[p])
			oldQ := find(nums, used, sortedNums[q])
			return []int{oldP, oldQ}
		} else if sum < target {
			p++
		} else {
			q--
		}
	}
	return []int{}
}
func find(nums []int, used []bool, value int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == value && used[i] == false {
			used[i] = true
			break
		}
		i++
	}
	return i
}

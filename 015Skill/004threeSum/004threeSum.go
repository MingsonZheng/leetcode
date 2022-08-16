package main

import (
	"fmt"
	"sort"
)

// 15. 三数之和

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] { // 避免 a 重复
			continue
		}
		p := i + 1
		q := n - 1
		for p < q {
			if p >= i+2 && nums[p] == nums[p-1] { // 避免 b 重复
				p++
				continue
			}
			if q <= n-2 && nums[q] == nums[q+1] { // 避免 c 重复
				q--
				continue
			}
			sum := nums[p] + nums[q]
			if sum == -1*nums[i] {
				resultItem := make([]int, 0)
				resultItem = append(resultItem, nums[i])
				resultItem = append(resultItem, nums[p])
				resultItem = append(resultItem, nums[q])
				result = append(result, resultItem)
				p++
				q--
			} else if sum < -1*nums[i] {
				p++
			} else {
				q--
			}
		}
	}
	return result
}

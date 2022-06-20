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
	hashMap := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		hashMap[nums[i]] = i
	}
	for i := 0; i < n; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		} // 避免 a 重复，1， 1， 3， ...
		for j := i + 1; j < n; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			} // 避免 b 重复，1，2，2，...
			target := -1 * (nums[i] + nums[j])
			if _, ok := hashMap[target]; !ok {
				continue
			}
			k := hashMap[target]
			if k > j {
				resultItem := make([]int, 0)
				resultItem = append(resultItem, nums[i], nums[j], nums[k])
				result = append(result, resultItem)
			}
		}
	}
	return result
}

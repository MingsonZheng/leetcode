package main

import "fmt"

// 78. 子集

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	backtrack(nums, 0, []int{})
	return result
}

func backtrack(nums []int, k int, path []int) {
	if k == len(nums) {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	backtrack(nums, k+1, path)
	path = append(path, nums[k])
	backtrack(nums, k+1, path)
	path = path[:len(path)-1]
}

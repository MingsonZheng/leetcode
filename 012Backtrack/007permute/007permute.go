package main

import "fmt"

// 46. 全排列

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	path := make([]int, 0)
	backtrack(nums, 0, path)
	return result
}

func backtrack(nums []int, k int, path []int) {
	if k == len(nums) {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	for i := 0; i < len(nums); i++ {
		if contain(path, nums[i]) {
			continue
		}
		path = append(path, nums[i])
		backtrack(nums, k+1, path)
		path = path[:len(path)-1]
	}
}

func contain(path []int, num int) bool {
	for _, p := range path {
		if p == num {
			return true
		}
	}
	return false
}

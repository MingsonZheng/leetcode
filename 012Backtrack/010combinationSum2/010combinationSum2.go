package main

import "fmt"

// 40. 组合总和 II

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	hashTable := make(map[int]int, 0)
	for i := 0; i < len(candidates); i++ {
		if _, ok := hashTable[candidates[i]]; !ok {
			hashTable[candidates[i]] = 1
		} else {
			hashTable[candidates[i]] = hashTable[candidates[i]] + 1
		}
	}
	nums := make([]int, 0)
	counts := make([]int, 0)
	for i := 0; i < len(candidates); i++ {
		if _, ok := hashTable[candidates[i]]; ok {
			nums = append(nums, candidates[i])
			counts = append(counts, hashTable[candidates[i]])
			delete(hashTable, candidates[i])
		}
	}
	backtrack(nums, counts, 0, target, []int{})
	return result
}

func backtrack(nums, counts []int, k, left int, path []int) {
	if left == 0 {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	if left < 0 || k == len(nums) {
		return
	}
	for count := 0; count <= counts[k]; count++ {
		for i := 0; i < count; i++ {
			path = append(path, nums[k])
		}
		backtrack(nums, counts, k+1, left-count*nums[k], path)
		for i := 0; i < count; i++ {
			path = path[:len(path)-1]
		}
	}
}

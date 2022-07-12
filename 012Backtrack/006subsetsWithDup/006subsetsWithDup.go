package main

import "fmt"

// 90. 子集 II

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

var result [][]int

func subsetsWithDup(nums []int) [][]int {
	result = make([][]int, 0)
	hm := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		count := 1
		if _, ok := hm[nums[i]]; ok {
			count += hm[nums[i]]
		}
		hm[nums[i]] = count
	}
	n := len(hm)
	uniqueNums := make([]int, n)
	counts := make([]int, n)
	k := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := hm[nums[i]]; ok {
			uniqueNums[k] = nums[i]
			counts[k] = hm[nums[i]]
			k++
			delete(hm, nums[i])
		}
	}
	backtrack(uniqueNums, counts, 0, []int{})
	return result
}

func backtrack(uniqueNums, counts []int, k int, path []int) {
	if k == len(uniqueNums) {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	// 有两个2，第一遍循环放一个，第二遍循环放两个，保证不重复
	for count := 0; count <= counts[k]; count++ {
		for i := 0; i < count; i++ {
			path = append(path, uniqueNums[k])
		}
		backtrack(uniqueNums, counts, k+1, path)
		for i := 0; i < count; i++ {
			path = path[:len(path)-1]
		}
	}
}

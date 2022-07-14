package main

import "fmt"

// 47. 全排列 II

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

var result [][]int

func permuteUnique(nums []int) [][]int {
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
	backtrack(uniqueNums, counts, 0, []int{}, len(nums))
	return result
}

func backtrack(uniqueNums, counts []int, k int, path []int, n int) {
	if k == n {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	for i := 0; i < len(uniqueNums); i++ {
		if counts[i] == 0 {
			continue
		}
		path = append(path, uniqueNums[i])
		counts[i]--
		backtrack(uniqueNums, counts, k+1, path, n)
		path = path[:len(path)-1]
		counts[i]++
	}
}

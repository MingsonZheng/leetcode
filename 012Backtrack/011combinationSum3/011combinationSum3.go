package main

import "fmt"

// 216. 组合总和 III

func main() {
	fmt.Println(combinationSum3(3, 7))
}

var result [][]int

func combinationSum3(k int, n int) [][]int {
	result = make([][]int, 0)
	backtrack(k, n, 1, 0, []int{})
	return result
}

func backtrack(k, n, step, sum int, path []int) {
	if sum == n && len(path) == k {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	if sum > n || len(path) > k || step > 9 {
		return
	}
	backtrack(k, n, step+1, sum, path)
	path = append(path, step)
	backtrack(k, n, step+1, sum+step, path)
	path = path[:len(path)-1]
}

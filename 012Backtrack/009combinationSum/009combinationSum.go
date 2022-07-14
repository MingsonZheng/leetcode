package main

import "fmt"

// 39. 组合总和

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	backtrack(candidates, 0, target, []int{})
	return result
}

func backtrack(candidates []int, k, left int, path []int) {
	if left == 0 {
		snapshot := make([]int, len(path))
		copy(snapshot, path)
		result = append(result, snapshot)
		return
	}
	if k == len(candidates) {
		return
	}
	for i := 0; i <= left/candidates[k]; i++ {
		for j := 0; j < i; j++ {
			path = append(path, candidates[k])
		}
		backtrack(candidates, k+1, left-i*candidates[k], path)
		for j := 0; j < i; j++ {
			path = path[:len(path)-1]
		}
	}
}

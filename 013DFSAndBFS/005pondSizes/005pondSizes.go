package main

import (
	"fmt"
	"sort"
)

// 面试题 16.19. 水域大小

func main() {
	land := [][]int{
		{0, 2, 1, 0},
		{0, 1, 0, 1},
		{1, 1, 0, 1},
		{0, 1, 0, 1},
	}
	fmt.Println(pondSizes(land))
}

var count int
var n int
var m int

func pondSizes(land [][]int) []int {
	n = len(land)
	m = len(land[0])
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] == 0 {
				count = 0
				dfs(land, i, j)
				result = append(result, count)
			}
		}
	}
	sort.Ints(result)
	return result
}

func dfs(land [][]int, curi, curj int) {
	count++
	land[curi][curj] = 1
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}}
	for d := 0; d < 8; d++ {
		newi := curi + dirs[d][0]
		newj := curj + dirs[d][1]
		if newi >= 0 && newi < n && newj >= 0 && newj < m &&
			land[newi][newj] == 0 {
			dfs(land, newi, newj)
		}
	}
}

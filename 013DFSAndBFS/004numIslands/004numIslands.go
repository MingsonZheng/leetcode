package main

import "fmt"

// 200. 岛屿数量

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}

var visited [][]bool
var h int
var w int

func numIslands(grid [][]byte) int {
	h = len(grid)
	w = len(grid[0])
	visited = make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	result := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if visited[i][j] != true && grid[i][j] == '1' {
				result++
				dfs(grid, i, j)
			}
		}
	}
	return result
}

func dfs(grid [][]byte, i, j int) {
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited[i][j] = true
	for k := 0; k < 4; k++ {
		newi := i + directions[k][0]
		newj := j + directions[k][1]
		if newi >= 0 && newi < h && newj >= 0 && newj < w &&
			visited[newi][newj] == false &&
			grid[newi][newj] == '1' {
			dfs(grid, newi, newj)
		}
	}
}

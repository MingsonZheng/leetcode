package main

import "fmt"

// 79. 单词搜索

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}

var existed bool
var h int
var w int

func exist(board [][]byte, word string) bool {
	h = len(board)
	w = len(board[0])
	existed = false
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			visited := make([][]bool, h)
			for k := 0; k < len(visited); k++ {
				visited[k] = make([]bool, w)
			}
			dfs(board, word, i, j, 0, visited)
			if existed {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int, visited [][]bool) {
	if existed == true {
		return
	}
	if word[k] != board[i][j] {
		return
	}
	visited[i][j] = true
	if k == len(word)-1 {
		existed = true
		return
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for di := 0; di < 4; di++ {
		nexti := i + directions[di][0]
		nextj := j + directions[di][1]
		if nexti >= 0 && nexti < h && nextj >= 0 && nextj < w &&
			!visited[nexti][nextj] {
			dfs(board, word, nexti, nextj, k+1, visited)
		}
	}
	visited[i][j] = false
}

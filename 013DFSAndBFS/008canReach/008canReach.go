package main

import "fmt"

// 1306. 跳跃游戏 III

func main() {
	arr := []int{4, 2, 3, 0, 3, 1, 2}
	fmt.Println(canReach(arr, 5))
}

var visited []bool
var reached bool

func canReach(arr []int, start int) bool {
	n := len(arr)
	visited = make([]bool, n)
	reached = false
	dfs(arr, start)
	return reached
}

func dfs(arr []int, curi int) {
	if reached {
		return
	}
	if arr[curi] == 0 {
		reached = true
		return
	}
	visited[curi] = true
	move2left := curi - arr[curi]
	if move2left >= 0 && move2left < len(arr) &&
		visited[move2left] == false {
		dfs(arr, move2left)
	}
	move2right := curi + arr[curi]
	if move2right >= 0 && move2right < len(arr) &&
		visited[move2right] == false {
		dfs(arr, move2right)
	}
}

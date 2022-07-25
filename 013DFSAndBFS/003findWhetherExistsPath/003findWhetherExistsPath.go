package main

import "fmt"

// 面试题 04.01. 节点间通路

func main() {
	graph := [][]int{
		{0, 1},
		{0, 2},
		{1, 2},
		{1, 2},
	}
	fmt.Println(findWhetherExistsPath(3, graph, 0, 2))

	graph = [][]int{
		{0, 2},
		{1, 3},
		{1, 4},
		{1, 3},
		{2, 3},
		{3, 4},
	}
	fmt.Println(findWhetherExistsPath(5, graph, 0, 4))
}

var visited []bool
var adj []map[int]bool
var found bool

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	visited = make([]bool, n)
	adj = make([]map[int]bool, n)
	found = false
	for i := 0; i < n; i++ {
		adj[i] = make(map[int]bool, 0)
	}
	for i := 0; i < len(graph); i++ {
		if !adj[graph[i][0]][graph[i][1]] {
			adj[graph[i][0]][graph[i][1]] = true
		}
	}
	dfs(start, target)
	return found
}

func dfs(cur, target int) {
	if found {
		return
	}
	if cur == target {
		found = true
		return
	}
	visited[cur] = true
	for next, _ := range adj[cur] {
		if !visited[next] {
			dfs(next, target)
		}
	}
}

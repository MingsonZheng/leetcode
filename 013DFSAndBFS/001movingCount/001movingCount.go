package main

import "fmt"

// 剑指 Offer 13. 机器人的运动范围

func main() {
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 1, 0))
}

var visited [][]bool
var count int

func movingCount(m int, n int, k int) int {
	visited = make([][]bool, m)
	count = 0
	for i, _ := range visited {
		visited[i] = make([]bool, n)
	}
	dfs(0, 0, m, n, k)
	return count
}

func dfs(i, j, m, n, k int) {
	visited[i][j] = true
	count++
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for di := 0; di < 4; di++ {
		newi := i + directions[di][0]
		newj := j + directions[di][1]
		if newi >= m || newi < 0 || newj >= n || newj < 0 ||
			visited[newi][newj] == true ||
			check(newi, newj, k) == false {
			continue
		}
		dfs(newi, newj, m, n, k)
	}
}

func check(i, j, k int) bool {
	sum := 0
	for i != 0 {
		sum += i % 10
		i /= 10
	}
	for j != 0 {
		sum += j % 10
		j /= 10
	}
	return sum <= k
}

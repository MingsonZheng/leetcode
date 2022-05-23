package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 5))
}

// 240. 搜索二维矩阵 II
func searchMatrix(matrix [][]int, target int) bool {
	h := len(matrix)
	w := len(matrix[0])
	i := 0
	j := w - 1
	// 根据 matrix[i][j] 跟 target 的大小关系，一层一层的剥离
	for i <= h-1 && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j--
			continue
		}
		if matrix[i][j] < target {
			i++
			continue
		}
	}
	return false
}

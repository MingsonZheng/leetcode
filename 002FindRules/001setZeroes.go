package main

import "fmt"

func main() {
	matrix := [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	for _, array := range matrix {
		for _, value := range array {
			fmt.Printf("%v，", value)
		}
		fmt.Println()
	}
}

// 面试题 01.08. 零矩阵
func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	// 哪一行，哪一列要清空为 0
	zeroRows := make([]bool, n)
	zeroColumns := make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zeroRows[i] = true
				zeroColumns[j] = true
			}
		}
	}
	// 设置为 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if zeroRows[i] || zeroColumns[j] {
				matrix[i][j] = 0
			}
		}
	}
}

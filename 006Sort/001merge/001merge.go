package main

import "fmt"

// 面试题 10.01. 合并排序的数组

func main() {
	A := []int{1, 2, 3, 0, 0, 0}
	B := []int{2, 5, 6}
	merge(A, 3, B, 3)
	fmt.Println(A)
}

func merge(A []int, m int, B []int, n int) {
	k := m + n - 1
	i := m - 1
	j := n - 1
	for i >= 0 && j >= 0 {
		if A[i] >= B[j] {
			A[k] = A[i]
			k--
			i--
		} else {
			A[k] = B[j]
			k--
			j--
		}
	}
	for i >= 0 {
		A[k] = A[i]
		k--
		i--
	}
	for j >= 0 {
		A[k] = B[j]
		k--
		j--
	}
}

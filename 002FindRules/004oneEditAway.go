package main

import (
	"fmt"
	"math"
)

func main() {
	first := "pale"
	second := "ple"
	fmt.Println(oneEditAway(first, second))
}

// 面试题 01.05. 一次编辑
func oneEditAway(first string, second string) bool {
	n := len(first)
	m := len(second)
	// 长度相差大于 1，无法通过一次编辑匹配
	if math.Abs(float64(n-m)) > 1 {
		return false
	}

	i := 0
	j := 0
	// 碰到第一个不能匹配的字符
	for i < n && j < m && first[i] == second[j] {
		i++
		j++
	}
	// 替换 abdf abcf
	if n == m {
		i++
		j++
	} else if n > m { // 删除或插入abdf abf
		i++
	} else { // m > n 删除或插入
		j++
	}
	// 继续考察后面的元素，必须完全匹配
	for i < n && j < m {
		if first[i] != second[j] {
			return false
		}
		i++
		j++
	}
	return true
}

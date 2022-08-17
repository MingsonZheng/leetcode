package main

import "fmt"

// 461. 汉明距离

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	r := x ^ y
	mask := 1
	count := 0
	for i := 0; i < 31; i++ {
		if r&mask != 0 {
			count++
		}
		mask *= 2
	}
	return count
}

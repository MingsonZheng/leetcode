package main

import (
	"fmt"
	"math"
	"sort"
)

// 面试题 16.06. 最小差

func main() {
	a := []int{1, 3, 15, 11, 2}
	b := []int{23, 127, 235, 19, 8}
	fmt.Println(smallestDifference(a, b))
}

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	n := len(a)
	m := len(b)
	var minRet int64 = math.MaxInt64
	i := 0
	j := 0
	for i < n && j < m {
		if a[i] >= b[j] {
			minRet = int64(math.Min(float64(minRet), float64(a[i]-b[j])))
			j++
		} else {
			minRet = int64(math.Min(float64(minRet), float64(b[j]-a[i])))
			i++
		}
	}
	return int(minRet)
}

package main

import (
	"fmt"
	"sort"
)

// 56. 合并区间

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(intervals))
	intervals = [][]int{
		{1, 4},
		{4, 5},
	}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	curLeft := intervals[0][0]
	curRight := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= curRight {
			if intervals[i][1] > curRight {
				curRight = intervals[i][1]
			}
		} else {
			result = append(result, []int{curLeft, curRight})
			curLeft = intervals[i][0]
			curRight = intervals[i][1]
		}
	}
	result = append(result, []int{curLeft, curRight})
	return result
}

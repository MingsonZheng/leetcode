package main

import (
	"fmt"
	"sort"
)

// 252. 会议室

func main() {
	intervals := [][]int{
		{0, 30},
		{5, 10},
		{15, 20},
	}
	fmt.Println(canAttendMeetings(intervals))
	intervals = [][]int{
		{7, 10},
		{2, 4},
	}
	fmt.Println(canAttendMeetings(intervals))
}

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

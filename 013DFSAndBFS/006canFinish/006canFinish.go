package main

import "fmt"

// 207. 课程表

func main() {
	prerequisites := [][]int{
		{1, 0},
	}
	fmt.Println(canFinish(2, prerequisites))
	prerequisites = [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(canFinish(2, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjs := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adjs[i] = make([]int, 0)
	}
	indegrees := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		adjs[prerequisites[i][1]] = append(adjs[prerequisites[i][1]], prerequisites[i][0])
		indegrees[prerequisites[i][0]]++
	}
	zeroInDegrees := make([]int, 0)
	for i := 0; i < len(indegrees); i++ {
		if indegrees[i] == 0 {
			zeroInDegrees = append(zeroInDegrees, i)
		}
	}
	zeroInDegreesCount := 0
	for len(zeroInDegrees) > 0 {
		coursei := zeroInDegrees[0]
		zeroInDegrees = zeroInDegrees[1:]
		zeroInDegreesCount++
		for _, coursej := range adjs[coursei] {
			indegrees[coursej]--
			if indegrees[coursej] == 0 {
				zeroInDegrees = append(zeroInDegrees, coursej)
			}
		}
	}
	return zeroInDegreesCount == numCourses
}

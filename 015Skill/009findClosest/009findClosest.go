package main

import (
	"fmt"
	"math"
)

// 面试题 17.11. 单词距离

func main() {
	words := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	word1 := "a"
	word2 := "student"
	fmt.Println(findClosest(words, word1, word2))
}

func findClosest(words []string, word1 string, word2 string) int {
	w1ps := make([]int, 0)
	w2ps := make([]int, 0)
	for i := 0; i < len(words); i++ {
		word := words[i]
		if word == word1 {
			w1ps = append(w1ps, i)
		} else if word == word2 {
			w2ps = append(w2ps, i)
		}
	}
	p1 := 0
	p2 := 0
	minRet := math.MaxInt32
	for p1 < len(w1ps) && p2 < len(w2ps) {
		pos1 := w1ps[p1]
		pos2 := w2ps[p2]
		if pos1 > pos2 {
			if minRet > pos1-pos2 {
				minRet = pos1 - pos2
			}
			p2++
		} else {
			if minRet > pos2-pos1 {
				minRet = pos2 - pos1
			}
			p1++
		}
	}
	return minRet
}

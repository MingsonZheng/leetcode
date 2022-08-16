package main

import (
	"fmt"
	"math"
)

// 76. 最小覆盖子串

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	minWSize := math.MaxInt32
	minWStart := -1
	minWEnd := -1
	tmap := make(map[byte]int, 0) // 模式串
	wmap := make(map[byte]int, 0) // 滑动窗口
	// 构建映射数组
	for i := 0; i < len(t); i++ {
		count := 1
		if value, ok := tmap[t[i]]; ok {
			count += value
		}
		tmap[t[i]] = count
	}
	n := len(s)
	l := 0
	r := -1
	for l < n && r < n {
		for !match(wmap, tmap) {
			r++
			if r > n-1 {
				break
			}
			c := s[r]
			if _, ok := tmap[c]; ok {
				count := 1
				if _, ok := wmap[c]; ok {
					count += wmap[c]
				}
				wmap[c] = count
			}
		}
		if match(wmap, tmap) {
			if minWSize > r-l+1 {
				minWSize = r - l + 1
				minWStart = l
				minWEnd = r
			}
			c := s[l]
			if _, ok := tmap[c]; ok {
				count := wmap[c]
				if count-1 == 0 {
					delete(wmap, c)
				} else {
					wmap[c] = count - 1
				}
			}
			l++
		}
	}
	if minWStart == -1 {
		return ""
	}
	return s[minWStart : minWEnd+1]
}

func match(wmap, tmap map[byte]int) bool {
	for key, value := range tmap {
		if _, ok := wmap[key]; !ok {
			return false
		}
		if wmap[key] < value {
			return false
		}
	}
	return true
}

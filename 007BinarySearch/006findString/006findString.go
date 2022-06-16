package main

import "fmt"

// 面试题 10.05. 稀疏数组搜索

func main() {
	words := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	s := "ta"
	//fmt.Println(findString(words, s))
	words = []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	s = "ball"
	fmt.Println(findString(words, s))
}

func findString(words []string, s string) int {
	low := 0
	high := len(words) - 1
	for low <= high {
		mid := low + (high-low)/2
		if words[mid] == s {
			return mid
		} else if words[mid] == "" {
			if words[low] == s {
				return low
			} else {
				low++
			}
		} else if words[mid] < s {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

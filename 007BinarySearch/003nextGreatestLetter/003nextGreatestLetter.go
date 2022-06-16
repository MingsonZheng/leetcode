package main

import "fmt"

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')
	fmt.Printf("%c", nextGreatestLetter(letters, target))
	fmt.Println()
	target = byte('c')
	fmt.Printf("%c", nextGreatestLetter(letters, target))
	fmt.Println()
	target = byte('d')
	fmt.Printf("%c", nextGreatestLetter(letters, target))
	fmt.Println()
}

func nextGreatestLetter(letters []byte, target byte) byte {
	low := 0
	high := len(letters) - 1
	for low <= high {
		mid := low + (high-low)/2
		c := letters[mid]
		if c > target {
			if mid == 0 || letters[mid-1] <= target {
				return letters[mid]
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return letters[0] // 这个题目的特殊要求
}

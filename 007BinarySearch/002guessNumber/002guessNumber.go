package main

import "fmt"

// 374. 猜数字大小

func main() {
	n := 10
	pick := 6
	fmt.Println(guessNumber(n, pick))
	n = 1
	pick = 1
	fmt.Println(guessNumber(n, pick))
	n = 2
	pick = 1
	fmt.Println(guessNumber(n, pick))
	n = 2
	pick = 2
	fmt.Println(guessNumber(n, pick))
}

func guessNumber(n int, pick int) int {
	low := 1
	high := n
	for low <= high {
		mid := low + (high-low)/2
		ret := guess(mid, pick)
		if ret == 0 {
			return mid
		} else if ret == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func guess(num int, pick int) int {
	if pick < num {
		return -1
	} else if pick > num {
		return 1
	} else {
		return 0
	}
}

package main

import "fmt"

// 231. 2 的幂

func main() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3))
}

func isPowerOfTwo(n int) bool {
	for n != 0 {
		if (n & 1) == 1 {
			if (n >> 1) == 0 {
				return true
			} else {
				return false
			}
		}
		n >>= 1
	}
	return false
}

package main

import "fmt"

// 191. 位1的个数
// 解法 2：自身

func main() {
	var num uint32
	num = 00000000000000000000000000001011
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	oneCount := 0
	for num != 0 {
		if num&1 == 1 {
			oneCount++
		}
		num >>= 1
	}
	return oneCount
}

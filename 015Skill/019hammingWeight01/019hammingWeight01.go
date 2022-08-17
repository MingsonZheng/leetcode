package main

import "fmt"

// 191. 位1的个数
// 解法 1：掩码

func main() {
	var num uint32
	num = 00000000000000000000000000001011
	fmt.Println(hammingWeight(num))
}

func hammingWeight(num uint32) int {
	oneCount := 0
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			oneCount++
		}
		mask <<= 1
	}
	return oneCount
}

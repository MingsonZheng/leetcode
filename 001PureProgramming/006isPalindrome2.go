package main

import "fmt"

func main() {
	fmt.Println(isPalindrome2(121))
	fmt.Println(isPalindrome3(121))
}

// 9. 回文数
func isPalindrome2(x int) bool {
	// -2147483648 ~ 2147483647
	digits := make([]int, 10)
	if x < 0 {
		return false
	}
	k := 0
	// 将 x 转化成字符串数组
	for x != 0 {
		digits[k] = x % 10
		x = x / 10
		k++
	}

	// 判断回文串
	for i := 0; i < k/2; i++ { // 举例验证
		if digits[i] != digits[k-i-1] { // 举例验证
			return false
		}
	}
	return true
}

func isPalindrome3(x int) bool {
	// -2147483648 ~ 2147483647
	if x < 0 {
		return false
	}
	backupX := x
	y := 0       // y 为 x 反转之后的值
	for x != 0 { // 将 x 转化成字符串数组的过程计算 y
		y = y*10 + x%10
		x = x / 10
	}
	return backupX == y
}

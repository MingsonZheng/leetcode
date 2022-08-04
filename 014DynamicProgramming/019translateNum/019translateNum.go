package main

import "fmt"

// 剑指 Offer 46. 把数字翻译成字符串

func main() {
	fmt.Println(translateNum(12258))
}

func translateNum(num int) int {
	if num <= 9 {
		return 1
	}
	// 把十进制转化成数字数组
	digitList := make([]int, 0)
	for num != 0 {
		digitList = append(digitList, num%10)
		num /= 10
	}
	n := len(digitList)
	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = digitList[n-i-1]
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if i-2 >= 0 && isValid(digits[i-2], digits[i-1]) {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

func isValid(a, b int) bool {
	if a == 1 {
		return true
	}
	if a == 2 && b >= 0 && b <= 5 {
		return true
	}
	return false
}

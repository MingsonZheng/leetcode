package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(strToInt(" -42"))
	fmt.Println(strToInt("4193abc"))
	fmt.Println(strToInt("abc978"))
	fmt.Println(strToInt("-9182838383838"))
	fmt.Println(strToInt("9182838383838"))
	fmt.Println(strToInt(""))
	fmt.Println(strToInt(" "))
}

// 剑指 Offer 67. 把字符串转换成整数
// " -42" 前导空格，符号位
// "4193abc" 后置非数字字符
// "abc978" 返回 0
// "-9182838383838" 超过整数范围 返回 math.MinInt32
// "9182838383838" 超过整数范围 返回 math.MaxInt32
// "" 空字符串 返回 0
// " " 全文空格 返回 0
func strToInt(str string) int {
	chars := []byte(str)
	n := len(chars)
	// 处理空
	if n == 0 {
		return 0
	}
	// 处理前置空格
	i := 0
	for i < n && chars[i] == ' ' {
		i++
	}
	// 全为空格
	if i == n {
		return 0
	}
	// 处理符号
	sign := 1
	c := chars[i]
	if c == '-' {
		sign = -1
		i++
	} else if c == '+' {
		sign = 1
		i++
	}
	// 真正处理数字
	// 整数范围 -2147483648 ~ 2147483647
	intAbsHigh := 214748364
	result := 0
	for i < n && chars[i] >= '0' && chars[i] <= '9' {
		d := chars[i] - '0'
		// 判断再乘以 10，加 d 之后，是否越界
		if result > intAbsHigh {
			if sign == 1 {
				return math.MaxInt32 // 214748365d
			} else {
				return math.MinInt32 // -214748365d
			}
		}
		if result == intAbsHigh {
			if (sign == 1) && (d > 7) {
				return math.MaxInt32 // 2147483648
			}
			if (sign == -1) && (d > 8) {
				return math.MinInt32 // -2147483649
			}
		}
		// 正常逻辑
		result = result*10 + int(d)
		i++
	}
	return sign * result
}

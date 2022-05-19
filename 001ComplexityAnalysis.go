package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr3("1.1.1.1"))
}

// 001复杂度分析

// 1. 两数之和
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 1108. IP 地址无效化
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

func defangIPaddr2(address string) string {
	sb := strings.Builder{}
	for i := 0; i < len(address); i++ {
		c := address[i]
		if c != '.' {
			sb.WriteByte(c)
		} else {
			sb.Write([]byte("[.]"))
		}
	}
	return sb.String()
}

// 用 []byte 实现：byte 是 uint8 的别名，占 8 位，所以处理 ASCII 没问题
func defangIPaddr3(address string) string {
	origin := []byte(address)
	n := len(origin)
	newN := n + 2*3
	newString := make([]byte, newN)
	k := 0
	for i := 0; i < n; i++ {
		if origin[i] != '.' {
			newString[k] = origin[i]
			k++
		} else {
			newString[k] = '[' // go 中 newString[K++] 编译不通过
			k++
			newString[k] = '.'
			k++
			newString[k] = ']'
			k++
		}
	}
	return string(newString)
}

// 用 []rune: rune 是 int32 的别名，占 32 位，4 个字节，如果处理中文字符更适合用 rune
func defangIPaddr4(address string) string {
	origin := []rune(address)
	n := len(origin)
	newN := n + 2*3
	newString := make([]rune, newN)
	k := 0
	for i := 0; i < n; i++ {
		if origin[i] != '.' {
			newString[k] = origin[i]
			k++
		} else {
			newString[k] = '[' // go 中 newString[K++] 编译不通过
			k++
			newString[k] = '.'
			k++
			newString[k] = ']'
			k++
		}
	}
	return string(newString)
}

func defangIPaddr666(address string) string {
	addr := strings.Split(address, ".")
	res := ""
	for i := 0; i < len(addr); i++ {
		res += addr[i]
		if i < len(addr)-1 {
			res += "[.]"
		}
	}
	return res
}

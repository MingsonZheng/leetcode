package main

import (
	"fmt"
	"math"
)

// 42. 接雨水
// 解法 3 单调栈解法
// 单调栈横向一层一层计算
// 栈里面存储位置一直往下的下标，一直往下，后面有往上的才有可能蓄水；一直往上不可能蓄水
// 遍历 height 元素 i
// 1）i <= top，入栈
// 2）i > top，出现凹槽，计算凹槽面积，计算之后出栈，相当于填平这一层的凹槽

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	n := len(height)
	result := 0
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, i) // 存下标
			continue
		}
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if height[top] == height[i] {
				stack = stack[:len(stack)-1]
				stack = append(stack, i)
				break
			} else if height[top] > height[i] {
				stack = append(stack, i)
				break
			} else { // 找到凹槽了
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					stack = append(stack, i)
					break
				}
				left := stack[len(stack)-1]
				h := int(math.Min(float64(height[left]), float64(height[i]))) - height[mid]
				w := i - left - 1
				result += h * w
			}
		}
	}
	return result
}

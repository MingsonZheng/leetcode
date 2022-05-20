package main

import "fmt"

func main() {
	solution := "RGBY"
	guess := "GGRR"
	fmt.Println(masterMind(solution, guess))
}

// 面试题 16.15. 珠玑妙算
func masterMind(solution string, guess string) []int {
	n := len(solution)
	hited := make([]bool, n) // guess 中哪些字符已经猜中了
	used := make([]bool, n)  // solution 中哪些字符已经被匹配用掉了
	// 先计算猜中的
	hitCount := 0
	for i := 0; i < n; i++ {
		if solution[i] == guess[i] {
			hited[i] = true
			used[i] = true
			hitCount++
		}
	}
	// 再计算伪猜中的
	fakeHitCount := 0
	for i := 0; i < n; i++ {
		if hited[i] {
			continue
		}
		// 拿每个 guess 中的字符到 solution 中查找
		for j := 0; j < n; j++ {
			if solution[j] == guess[i] && !used[j] {
				used[j] = true
				fakeHitCount++
				break
			}
		}
	}
	return []int{hitCount, fakeHitCount}
}

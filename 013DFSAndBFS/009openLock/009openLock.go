package main

import "fmt"

// 752. 打开转盘锁

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}

func openLock(deadends []string, target string) int {
	deadSet := make(map[string]bool, 0)
	for _, d := range deadends {
		deadSet[d] = true
	}
	if deadSet["0000"] {
		return -1
	}
	queue := make([]string, 0)
	visited := make(map[string]bool, 0)
	queue = append(queue, "0000")
	visited["0000"] = true
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		k := 0
		for k < size {
			node := queue[0]
			queue = queue[1:]
			k++
			if node == target {
				return depth
			}
			newNodes := genNewNode(node)
			for _, newNodes := range newNodes {
				if visited[newNodes] || deadSet[newNodes] {
					continue
				}
				queue = append(queue, newNodes)
				visited[newNodes] = true
			}
		}
		depth++
	}
	return -1
}

func genNewNode(node string) []string {
	newNodes := make([]string, 0)
	change := []int{-1, 1}
	for i := 0; i < 4; i++ {
		for k := 0; k < 2; k++ {
			newNode := make([]byte, 4)
			for j := 0; j < i; j++ {
				newNode[j] = node[j]
			}
			for j := i + 1; j < 4; j++ {
				newNode[j] = node[j]
			}
			newC := fmt.Sprintf("%d", (int(node[i]-'0')+change[k]+10)%10)
			newNode[i] = newC[0]
			newNodes = append(newNodes, string(newNode))
		}
	}
	return newNodes
}

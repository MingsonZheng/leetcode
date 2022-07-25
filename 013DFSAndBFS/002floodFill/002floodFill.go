package main

import "fmt"

// 面试题 08.10. 颜色填充

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	floodFill(image, 1, 1, 2)
	fmt.Println(image)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	n := len(image)
	m := len(image[0])
	dfs(image, n, m, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(image [][]int, n, m, sr, sc, color, newColor int) {
	image[sr][sc] = newColor
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for k := 0; k < 4; k++ {
		newr := sr + dirs[k][0]
		newc := sc + dirs[k][1]
		if newr < 0 || newr >= n || newc < 0 || newc >= m ||
			image[newr][newc] != color ||
			image[newr][newc] == newColor {
			continue
		}
		dfs(image, n, m, newr, newc, color, newColor)
	}
}

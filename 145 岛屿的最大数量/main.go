package main

import (
	"fmt"
	"math"
)

func maxAreaOfIsland(grid [][]int) int {
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 0
		sum := 0
		for z := 0; z < 4; z++ {
			newX, newY := z+dir[i][0], j+dir[i][1]
			sum += dfs(newX, newY)
		}
		return 1 + sum
	}
	ans := math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ret := maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}})
	fmt.Println(ret)
}

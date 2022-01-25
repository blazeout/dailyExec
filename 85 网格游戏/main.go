package main

import (
	"fmt"
	"math"
)

// 力扣第2017道题
// 利用前缀数组来求解
func gridGame(grid [][]int) int64 {
	n := len(grid[0])
	// 构建一个前缀和二维数组来求解
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i-1]
		dp[1][i] = dp[1][i-1] + grid[1][i-1]
	}
	ans := math.MaxInt64
	for i := 1; i <= n; i++ {
		ans = min(ans, max(dp[0][n]-dp[0][i], dp[1][i-1]))
	}
	return int64(ans)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func main() {
	grid := [][]int{{3, 3, 1}, {8, 5, 2}}
	fmt.Println(gridGame(grid))
	fmt.Println(grid)
}

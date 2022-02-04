package _02_黄金矿工

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxRes := 0
	isvisited := make([][]bool, m)
	for i := range isvisited {
		isvisited[i] = make([]bool, n)
	}
	// dfs函数返回上下左右四个方向的值， 然后再从四个值中选出最大值
	var dfs func(i, j int, sum int) int
	dfs = func(i, j int, sum int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || isvisited[i][j] {
			return 0
		}
		isvisited[i][j] = true
		ret1 := dfs(i, j+1, sum)
		ret2 := dfs(i, j-1, sum)
		ret3 := dfs(i+1, j, sum)
		ret4 := dfs(i-1, j, sum)
		isvisited[i][j] = false
		ret1 = max(ret1, ret2)
		ret3 = max(ret3, ret4)
		ret1 = max(ret1, ret3)
		sum += ret1
		return sum + grid[i][j]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				maxRes = max(maxRes, dfs(i, j, 0))
			}
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

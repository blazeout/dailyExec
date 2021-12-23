package _1_岛屿的周长

func islandPerimeter(grid [][]int) int {
	// for 循环遍历grid数组, 如果是1, 那么就判断四周是否是陆地, 如果是陆地那么就减1
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += dfs(i-1, j, grid) + dfs(i+1, j, grid) + dfs(i, j-1, grid) + dfs(i, j+1, grid)
			}
		}
	}
	return res
}

func dfs(i, j int, grid [][]int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 1
	}
	if grid[i][j] == 1 {
		return 0
	}
	return 1
}

package main

type pair struct {
	x int
	y int
}

func numEnclaves(grid [][]int) int {
	res := 0
	m, n := len(grid), len(grid[0])
	//if m == 1 || n == 1 {
	//	return 0
	//}
	dir := []pair{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	// dfs的作用是把边界上面的1和能达到的1全部改为2
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		for _, v := range dir {
			newX, newY := i+v.x, j+v.y
			dfs(newX, newY)
		}
	}
	// 两个横向
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	// 两个纵向
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 1 {
				res += 1
			}
		}
	}
	return res
}

func main() {
	numEnclaves([][]int{{0, 0, 1, 1, 0}})
}

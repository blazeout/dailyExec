package main

import (
	"fmt"
)

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var tmp int
			fmt.Scan(&tmp)
			matrix[i][j] = tmp
		}
	}
	res := 0
	isVisited := make([][]int, n)
	for i := range isVisited {
		isVisited[i] = make([]int, m)
	}
	var dfs func(i, j int, num int)
	dfs = func(i, j int, num int) {
		if i < 0 || j < 0 || i >= n || j >= m || matrix[i][j] != num || isVisited[i][j] != 0 {
			return
		}
		isVisited[i][j] = 1
		matrix[i][j] = 4
		dfs(i+1, j, num)
		dfs(i-1, j, num)
		dfs(i, j+1, num)
		dfs(i, j-1, num)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != 4 {
				dfs(i, j, matrix[i][j])
				res++
			}
		}
	}
	fmt.Println(res)
}

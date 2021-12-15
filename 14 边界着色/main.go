package main

import "fmt"

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	n := len(grid)
	m := len(grid[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	revise := make([][]int, n)
	for i := range revise {
		revise[i] = make([]int, m)
	}
	for i := range grid {
		copy(revise[i], grid[i])
	}
	var dfs func(i, j, pre int)
	dfs = func(i, j, pre int) {
		// dfs 作用是判断, 将联通分量的点全部标记一遍
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] != pre || visited[i][j] {
			return
		}
		visited[i][j] = true
		dfs(i+1, j, pre)
		dfs(i, j+1, pre)
		dfs(i, j-1, pre)
		dfs(i-1, j, pre)
	}
	dfs(row, col, grid[row][col])
	// 找到这个点, 然后dfs遍历, 将属于连通分量的点全部给标记一遍

	for i, v := range visited {
		for j, isvisited := range v {
			// 代表这个点属于连通分量
			if isvisited == true {
				if i == 0 || i == n-1 || j == 0 || j == m-1{
					revise[i][j] = color
					continue
				}
				// 这里就判断周围4个方向是否有不同颜色的格子, 如果有就属于边界
				if isBoard(i, j, grid, grid[i][j]) {
					revise[i][j] = color
				}
			}
		}
	}
	return revise
}
func isBoard2(i, j int, grid [][]int, color int) bool {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1  {
		return false // 超出边界
	}
	if grid[i][j] != color {
		return true
	}
	return false
}
func isBoard(i, j int, grid [][]int, color int) bool {
	if isBoard2(i+1,j, grid, color) || isBoard2(i, j+1, grid, color) || isBoard2(i-1, j, grid, color) || isBoard2(i, j-1, grid, color){
		return true
	}
	return false
}

func main() {

	grid := make([][]int, 3)
	for i := range grid {
		grid[i] = make([]int, 3)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = 1
		}
	}
	row := 1
	col := 1
	color := 2
	revise := colorBorder(grid, row, col, color)
	fmt.Println(revise)
	//slice1 := make([]int, 0, 4)
	//slice1 = append(slice1, 1, 2, 3)
	//
	//slice2 := slice1[:len(slice1)-1]
	//slice2 = append(slice2,11,12,13,14,15)
	//slice2[0] = 10
	//
	//fmt.Println(slice1)
	//fmt.Println(slice2)
}

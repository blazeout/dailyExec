package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	// 定义一个反转变量, 偶数向上,奇数向下
	reverse := 0
	i, j := 0, 0
	m, n := len(mat), len(mat[0])
	maxCountOfReverse := m + n - 1
	result := make([]int, 0)
	for i < m && j < n && reverse <= maxCountOfReverse {
		result = append(result, mat[i][j])
		if reverse%2 == 0 {
			// 偶数往上走
			if i-1 >= 0 && j+1 < n {
				i -= 1
				j += 1
			} else {
				// 右上方向是没有右边界 才往下拐
				if j+1 < n {
					j += 1
				} else {
					i += 1
				}
				reverse++
			}
		} else {
			// 奇数往左下走 i+1 j-1
			if i+1 < m && j-1 >= 0 {
				i += 1
				j -= 1
			} else {
				// 左下方向是没有下边界才往右拐
				if i+1 < m {
					i += 1
				} else {
					j += 1
				}
				reverse++
			}
		}
	}
	return result
}

func main() {
	arr := [][]int{{1, 2}, {3, 4}}
	diagonalOrder := findDiagonalOrder(arr)
	fmt.Println(diagonalOrder)
}

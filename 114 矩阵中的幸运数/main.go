package main

import "fmt"

func luckyNumbers(matrix [][]int) []int {
	// 构建两个数组即可, 一个最小数组, 一个最大数组
	minOfRow := make([]int, len(matrix))
	maxOfCol := make([]int, len(matrix[0]))
	m, n := len(minOfRow), len(maxOfCol)
	for i := 0; i < m; i++ {
		minValue := matrix[i][0]
		for j := 1; j < n; j++ {
			if matrix[i][j] < minValue {
				minValue = matrix[i][j]
			}
		}
		minOfRow[i] = minValue
	}
	for j := 0; j < n; j++ {
		maxValue := matrix[0][j]
		for i := 1; i < m; i++ {
			if matrix[i][j] > maxValue {
				maxValue = matrix[i][j]
			}
		}
		maxOfCol[j] = maxValue
	}

	res := []int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == minOfRow[i] && matrix[i][j] == maxOfCol[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}

func main() {
	ret := luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}})
	fmt.Println(ret)
}

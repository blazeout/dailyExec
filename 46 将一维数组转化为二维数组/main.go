package _6_将一维数组转化为二维数组

func construct2DArray1(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	res := [][]int{}
	for i := 0; i < len(original); i += n {
		res = append(res, original[i:i+n])
	}
	return res
}

func construct2DArray2(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = original[i*n+j]
		}
	}
	return res
}

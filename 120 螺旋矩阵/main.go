package _20_螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		// 首先遍历top行
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++
		if top > bottom {
			break
		}
		// 然后遍历right列
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for j := right; j >= left; j-- {
			res = append(res, matrix[bottom][j])
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

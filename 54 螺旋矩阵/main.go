package _4_螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := []int{}
	// 控制的
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		// 先打印最上面一行
		for j := left; j <= right; j++ {
			res = append(res, matrix[top][j])
		}
		top++
		if top > bottom {
			break
		}
		// 在打印最右边的一行
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

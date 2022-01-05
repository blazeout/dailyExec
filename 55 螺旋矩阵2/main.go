package _5_螺旋矩阵2

// 自己的解法
func generateMatrix1(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	count := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for {
		for j := left; j <= right; j++ {
			res[top][j] = count
			count += 1
		}
		top++
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res[i][right] = count
			count += 1
		}
		right--
		if right < left {
			break
		}
		for j := right; j >= left; j-- {
			res[bottom][j] = count
			count += 1
		}
		bottom--
		if bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			res[i][left] = count
			count++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}

// K神的解法
func generateMatrix2(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	l, r, t, b := 0, n-1, 0, n-1
	tar := n * n
	num := 1
	for num <= tar {
		for j := l; j <= r; j++ {
			matrix[t][j] = num
			num++
		}
		t++
		for i := t; i <= b; i++ {
			matrix[i][r] = num
			num++
		}
		r--
		for j := r; j >= l; j-- {
			matrix[b][j] = num
			num++
		}
		b--
		for i := b; i >= t; i-- {
			matrix[i][l] = num
			num++
		}
		l++
	}
	return matrix
}

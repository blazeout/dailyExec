package _28_Z字形变化

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	// 1列有numRows个数, 旁边上升的数是numRows-2, 总的一个周期就是num + num-2 = 2num-2
	// 一个周期的列数就是, 1 + r-2 = r-1列
	// 然后我们就总共需要 c := (len(s)/t + 1 ) * (r-1)
	n, r := len(s), numRows
	// t是一个周期内的字符数量
	t := r + r - 2
	c := (n/t + 1) * (r - 1)
	matrix := make([][]byte, numRows)
	for i := range matrix {
		matrix[i] = make([]byte, c)
	}
	x, y := 0, 0
	for i := 0; i < len(s); i++ {
		matrix[x][y] = s[i]
		// 重点 i % t < r - 1 时就应该往下走, 大于时就应该往右上走. 都是一半一半
		if i%t < r-1 {
			// 往下走
			x++
		} else {
			x--
			y++
		}
	}
	ans := make([]byte, n)
	index := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] > 0 {
				ans[index] = matrix[i][j]
				index++
			}
		}
	}
	return string(ans)
}

func convert2(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	// 1列有numRows个数, 旁边上升的数是numRows-2, 总的一个周期就是num + num-2 = 2num-2
	// 一个周期的列数就是, 1 + r-2 = r-1列
	// 然后我们就总共需要 c := (len(s)/t + 1 ) * (r-1)
	n, r := len(s), numRows
	// t是一个周期内的字符数量
	t := r + r - 2
	// c := (n/t + 1) * (r-1)
	// 注意到原始矩阵有很多空余的地方, 我们只需要将这一行的数据添加到matrix的某一行结尾即可
	matrix := make([][]byte, r)
	x := 0
	for i := 0; i < n; i++ {
		matrix[x] = append(matrix[x], s[i])
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	ans := []byte{}
	for i := 0; i < len(matrix); i++ {
		ans = append(ans, matrix[i]...)
	}
	return string(ans)
}

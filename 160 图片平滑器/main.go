package _60_图片平滑器

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	var average func(i, j int) (int, int)
	average = func(i, j int) (int, int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0, 0
		}
		return img[i][j], 1
	}
	var get func(i, j int) (int, int)
	get = func(i, j int) (int, int) {
		sum, count := 0, 0
		for z := i - 1; z <= i+1; z++ {
			for x := j - 1; x <= j+1; x++ {
				tmp1, tmp2 := average(z, x)
				sum += tmp1
				count += tmp2
			}
		}
		return sum, count
	}
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			sum, count := get(i, j)
			matrix[i][j] = sum / count
		}
	}
	return matrix
}

package _67_自除数

func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		if check(i) {
			res = append(res, i)
		}
	}
	return res
}

func check(i int) bool {
	temp := i
	for temp > 0 {
		remain := temp % 10
		if remain == 0 || i%remain != 0 {
			return false
		}
		temp /= 10
	}
	return true
}

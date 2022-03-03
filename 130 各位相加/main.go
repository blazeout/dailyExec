package _30_各位相加

func addDigits(num int) int {
	for num > 9 {
		num = ret(num)
	}
	return num
}

func ret(a int) int {
	sum := 0
	for a > 0 {
		sum += a % 10
		a /= 10
	}
	return sum
}

// 一步到位, 数根原理
func addDigits2(num int) int {
	return (num-1)%9 + 1
}

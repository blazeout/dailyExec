package _7_将数字变成0的操作

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num&1 == 0 {
			// 偶数
			num >>= 1
		} else {
			num -= 1
		}
		count++
	}
	return count
}

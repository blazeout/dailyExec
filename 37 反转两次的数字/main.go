package _7_反转两次的数字

// 如果为0直接返回true, 然后判断最后一个数字是否为0. 如果为0那么就不相等, 不为0就相等
func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}
	if num%10 != 0 {
		return true
	}
	return false
}

package _6_有效的括号

func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}

	// 注：由于这里 len(s) 是偶数，所以循环结束后 x 也是偶数（这意味着可以通过配对来让括号平衡度为 0），无需判断 x 是否为奇数
	// x 是偶数是因为每遍历一个字符必然会改变 x 的奇偶性，而 x 的奇偶性在发生偶数次变化后的结果是 x 的奇偶性不变
	x := 0
	for i, ch := range s {
		if ch == '(' || locked[i] == '0' {
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}

	x = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			x++
		} else if x > 0 {
			x--
		} else {
			return false
		}
	}
	return true
}

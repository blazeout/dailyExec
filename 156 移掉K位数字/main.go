package _56_移掉K位数字

func removeKdigits(num string, k int) string {
	// 像 1432219 这样高位是递减的数字, 那么我们肯定想删除高位, 让低位顶上去
	// 像 1235311 这样高位是递增的数字, 那么我们肯定想保持高位是最小的, 所以我们要删除低位
	// 综合起来就是, 如果此时num[i] < 栈顶元素, 就让栈顶元素出栈, num[i]进栈
	stack := []byte{}
	for i := 0; i < len(num); i++ {
		// 我们不想让前导0入栈, 那么就是如果num[i] == '0' && len(stack) == 0 的情况我们就不入栈
		// 取反就是入栈
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		if num[i] != '0' || len(stack) > 0 {
			stack = append(stack, num[i])
		}
	}
	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

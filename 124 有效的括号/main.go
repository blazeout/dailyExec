package _24_有效的括号

// 有效的括号
// 利用栈的先进后出特点来做

func isValid(s string) bool {
	if len(s) == 1 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			// 就是']', ')', '}'
			if len(stack) != 0 {
				if s[i] == ')' {
					if stack[len(stack)-1] == '(' {
						stack = stack[:len(stack)-1]
					} else {
						return false
					}
				} else if s[i] == ']' {
					if stack[len(stack)-1] == '[' {
						stack = stack[:len(stack)-1]
					} else {
						return false
					}
				} else {
					if stack[len(stack)-1] == '{' {
						stack = stack[:len(stack)-1]
					} else {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

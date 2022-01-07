package _7_括号的最大嵌套深度

// 自己实现的 O(n) O(n)
func maxDepth1(s string) int {
	// 用栈
	// 1. 如果遇到'(' 就放入栈中
	// 2. 如果遇到')' 就从栈顶取出一个'(' 并且计算当前的最大值 ret = 1 + len(stack)
	// 3. maxRes = max(maxRes, ret)
	stack := []byte{}
	maxRes := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				ret := 1 + len(stack)
				maxRes = max(maxRes, ret)
			}
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth2(s string) int {
	// 遇到'(' 就+1
	// 遇到')' 就-1
	ans := 0
	size := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			size++
			if size > ans {
				ans = size
			}
		} else if s[i] == ')' {
			size--
		}
	}
	return ans
}

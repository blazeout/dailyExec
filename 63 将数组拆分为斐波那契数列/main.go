package _3_将数组拆分为斐波那契数列

import "math"

func splitIntoFibonacci(num string) (F []int) {
	n := len(num)
	var dfs func(index, sum, prev int) bool
	dfs = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}
		cur := 0 // 当前拆分的数值总和
		// for 循环控制此层的 比如124557 第一层1 或者是 12
		for i := index; i < n; i++ {
			// 每个块的数字一定不要以0开头, 除非就是0本身
			if i > index && num[index] == '0' {
				break
			}
			cur = cur*10 + int(num[i]-'0')
			if cur > math.MaxInt32 {
				break
			}
			// 从第三个开始就需要满足:  F[i-2] + F[i-1] == F[i]
			if len(F) >= 2 {
				if cur < sum {
					// 此时还小于, 所以就对于目前的一段字符继续拼接
					continue
				}
				if cur > sum {
					// 此时已经大于前两个之和了, 所以也不需要拼接下面的字符了 直接break
					break
				}
			}
			// 进过上方两个条件的判断, 此时就是 cur == sum 或者 是len(F) < 2
			F = append(F, cur)
			// dfs 深入进入下一个拼接的字符串
			if dfs(i+1, prev+cur, cur) {
				return true
			}
			// 如果上方return false, 那么说明此路不通 不要继续下去了, 需要剪枝上一个添加到数组中的值
			F = F[:len(F)-1]
		}
		// 代表此路不通, 不能组成有效的斐波那契数列
		return false
	}
	dfs(0, 0, 0)
	return
}

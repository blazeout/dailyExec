package _1_Pow函数

// 快速幂算法 ?

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, n)
}

func quickMul(x float64, n int) float64 {
	// 迭代贡献值
	ans := 1.0
	xContribution := x
	for n > 0 {
		// 二进制位上的1
		if n & 1 == 1 {
			ans *= xContribution
		}
		// 每次循环更新贡献值
		xContribution *= xContribution
		n >>= 1
	}
	return ans
}

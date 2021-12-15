package _2_超级次方

const mod = 1337
func superPow(a int, b []int) int {
	// 倒序遍历 a^n =
	ans := 1
	for i := len(b)-1; i >= 0; i-- {
		ans = ans * pow(a, b[i]) % mod
		a = pow(a, 10)
	}
	return ans
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n & 1 == 1 {
			res = res*x % mod
		}
		x = x * x % mod
	}
	return res
}
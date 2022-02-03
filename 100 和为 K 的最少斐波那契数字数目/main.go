package _00_和为_K_的最少斐波那契数字数目

// 贪心, 每次找到一个距离K最近的斐波那契数字 然后减去
func findMinFibonacciNumbers(k int) (ans int) {
	f := []int{1, 1}
	for f[len(f)-1] < k {
		f = append(f, f[len(f)-1]+f[len(f)-2])
	}
	for i := len(f) - 1; k > 0; i-- {
		if k >= f[i] {
			k -= f[i]
			ans++
		}
	}
	return
}

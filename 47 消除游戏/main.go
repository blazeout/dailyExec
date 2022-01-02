package _7_消除游戏

func lastRemaining(n int) int {
	head := 1    // 最后剩的头部
	step := 1    // 步数
	left := true // 代表是否从左边开始
	for n > 1 {
		if left || n%2 != 0 { // 从左边开始移除, 或者从右边开始移除 但是要满足 长度 % 2 != 0 就满足需要移除
			head += step
		}
		n >>= 1
		step <<= 1
		left = !left
	}
	return head
}

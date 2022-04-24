package _93_二进制间距

func binaryGap(n int) int {
	// 先 生成一个32位的数组
	// arr := make([]int, 0)
	// for n > 0 {
	//     arr = append(arr, n & 1)
	//     n >>= 1
	// }
	maxRes := 0
	before := -1
	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			if before != -1 {
				maxRes = max(maxRes, i-before)
			}
			before = i
		}
		n >>= 1
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

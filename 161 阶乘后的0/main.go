package main

import "fmt"

// 将n 分解为 2 * 5, 2的个数肯定比5多, 所以我们只需要计算5的个数即可
func trailingZeroes(n int) int {
	count := 0
	for n > 0 {
		count += n / 5
		n = n / 5
	}
	fmt.Println(count)
	return count
}

func main() {
	trailingZeroes(5)
}

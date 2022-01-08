package main

import "fmt"

func grayCode(n int) []int {
	// 第一个整数是0
	// 一个整数在序列中出现不超过一次
	// 每队 相邻 的整数的二进制表示恰好一位不同
	// i ^ i >> 1
	// num 和 num右移1位做异或
	ans := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		ans[i] = i ^ (i >> 1)
	}
	return ans
}

func main() {
	ret := grayCode(2)
	fmt.Println(ret)
}

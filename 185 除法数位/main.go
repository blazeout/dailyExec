package main

import "fmt"

// 小明在尝试把一些分子为1的分数（1/x）转化为小数。
//使用普通计算器的除法功能可以实现，但是保留的小数位数非常有限。而小明希望得到n位小数，而且要从小数点后面第d位开始的n位小数。
//例如，x=13, d=4, n=3时，1/13=0.07692307692……，从小数点后第四位开始取三位数，答案是923。
//现在小明想要计算一些数更大的情况（2 <= x <= 10000, 1 <= d <= 1000000000, 1 <= n <= 10000），请你写个程序帮帮他。

func main() {
	d, n := 0, 0
	var x int
	fmt.Scan(&x, &d, &n)
	arr1 := make([]int, d+n)
	value := 1
	for i := 0; i < d+n; i++ {
		if value < x {
			arr1[i] = 0
			value *= 10
		} else {
			arr1[i] = value / x
			value = (value % x) * 10
		}
	}
	arr := make([]int, 0)
	for i := len(arr1) - 1; i >= 0 && n > 0; i-- {
		arr = append(arr, arr1[i])
		n--
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		sum *= 10
	}
	fmt.Println(sum / 10)
}

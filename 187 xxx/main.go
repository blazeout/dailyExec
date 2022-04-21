package main

import (
	"fmt"
)

func calc(value int) int {
	sum := 0
	for value > 0 {
		if value%10 == 0 {
			sum++
		}
		value /= 10
	}
	return sum
}

func main() {
	mpValue := map[int]int{}
	mpPath := map[int][]int{}
	n, m := 0, 0
	fmt.Scan(&n, &m)
	couse := 1
	arr := make([]int, n)
	// 按行读取值
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n; i++ {
		mpValue[couse] = arr[i]
		couse++
	}
	for j := 0; j < m; j++ {
		from, to := 0, 0
		fmt.Scan(&from, &to)
		mpPath[from] = append(mpPath[from], to)
	}
	maxRes := 0
	for from, to := range mpPath {
		for i := 0; i < len(to); i++ {
			temp := calc(mpValue[from] * mpValue[to[i]])
			if temp > maxRes {
				maxRes = temp
			}
		}
	}
	fmt.Println(maxRes)
}

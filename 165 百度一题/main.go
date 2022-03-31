package main

import (
	"fmt"
)

func reverse(num int) int {
	sum := 0
	for num > 0 {
		sum = num%10 + sum*10
		num /= 10
	}
	return sum
}

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	tmp := n
	maxRes := 0
	for i := 0; i < k; i++ {
		ans := reverse(tmp)
		if ans > maxRes {
			maxRes = ans
		}
		tmp += n
	}
	fmt.Println(maxRes)
}

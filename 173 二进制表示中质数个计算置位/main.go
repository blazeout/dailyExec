package main

import (
	"fmt"
)

func countPrimeSetBits(left int, right int) int {
	// 质数
	ans := 0
	mp := map[int]bool{}
	for i := left; i <= right; i++ {
		count := 0
		tmp := i
		for tmp > 0 {
			tmp &= tmp - 1
			count++
		}
		if mp[count] || isPrime(count) {
			mp[count] = true
			ans++
		}
	}
	return ans
}

func isPrime(a int) bool {

	if a <= 3 {
		return a > 1
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	ret := countPrimeSetBits(10, 15)
	fmt.Println(ret)
}

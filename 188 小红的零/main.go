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
	m := 0
	fmt.Scan(&m)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&arr[i])
	}
	sum := 0
	mp := map[int]int{}
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			if j == i {
				if _, ok := mp[arr[i]]; ok {
					sum += mp[arr[i]]
				} else {
					temp := calc(arr[i])
					sum += temp
					mp[arr[i]] = temp
				}
				continue
			}
			chenji := 1
			for z := i; z <= j; z++ {
				chenji *= arr[z]
			}
			if _, ok := mp[chenji]; ok {
				sum += mp[arr[i]*arr[j]]
			} else {
				temp := calc(chenji)
				sum += temp
				mp[chenji] = temp
			}
		}
	}
	fmt.Println(sum)
}

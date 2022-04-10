package main

import (
	"fmt"
	"sort"
)

func largestInteger(num int) int {
	// sort
	arr1, arr2 := []int{}, []int{}
	compare := []int{}
	tmp := num
	for tmp > 0 {
		x := tmp % 10
		compare = append(compare, x)
		if x%2 == 0 {
			// 偶数数组
			arr1 = append(arr1, x)
		} else {
			// 奇数数组
			arr2 = append(arr2, x)
		}
		tmp /= 10
	}
	for i, j := 0, len(compare)-1; i < j; i, j = i+1, j-1 {
		compare[i], compare[j] = compare[j], compare[i]
	}
	sort.Slice(arr1, func(i, j int) bool { return arr1[i] > arr1[j] })
	sort.Slice(arr2, func(i, j int) bool { return arr2[i] > arr2[j] })
	res := []int{}
	m, n := len(arr1), len(arr2)
	i, j := 0, 0
	count := 0
	for i < m || j < n {
		if compare[count]%2 == 0 {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
		count++
	}
	return calc(res)
}

func calc(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		sum *= 10
	}
	return sum / 10
}

func main() {
	ret := largestInteger(1234)
	fmt.Println(ret)
}

package main

import (
	"sort"
	"strconv"
)

func smallestNumber(num int64) int64 {
	flag := false
	if num < 0 {
		flag = true
		num = -num
	}
	arr := []int64{}
	for num != 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	if flag {
		sort.Slice(arr, func(i, j int) bool {
			// 递减排序
			return arr[i] > arr[j]
		})
	} else {
		sort.Slice(arr, func(i, j int) bool {
			// 递增排序
			return arr[i] < arr[j]
		})
	}
	str := ""
	if !flag && contains(0, arr) {
		// 包含0
		for _, v := range arr {
			if v != 0 {
				str += strconv.Itoa(int(v))
			}
		}
		str = str[0:1] + "0" + str[1:]
	} else if !flag && !contains(0, arr) {
		// 不包含0
		for _, v := range arr {
			str += strconv.Itoa(int(v))
		}
	} else {
		// 负数
		str += "-"
		for _, v := range arr {
			str += strconv.Itoa(int(v))
		}
	}
	v, _ := strconv.Atoi(str)
	return int64(v)
}

func contains(x int64, arr []int64) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return true
		}
	}
	return false
}

func main() {
	println(smallestNumber(-7460))
}

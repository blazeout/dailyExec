package main

import "sort"

func numFriendRequests(ages []int) int {
	// 1. 年龄小的不会向年龄大的发送好友请求
	// 2. 0.5 * age[x] + 7 < age[y] <= age[x]
	// 2.1 当age[x] <= 14 时不满足条件, 所以只需要看age[x] > 14的情况
	// 2.2 在遍历age[x]的同时, 使用双指针left 和 right 区间来规范y的区间代表x会给y发送的区间
	// 2.3 当age[x+1]时, 如果是左边不满足 0.5 * age[x] + 7 < age[y], 那么left++
	//     当age[x+1]时, 如果是右边不满足 right++
	left := 0
	right := 0
	sort.Ints(ages)
	count := 0
	for i := 0; i < len(ages); i++ {
		if ages[i] <= 14 {
			continue
		}
		for ages[left]*2 <= ages[i]+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= ages[i] {
			right++
		}
		count += right - left
	}
	return count
}

func main() {
	ages := []int{10, 20, 30, 100, 110, 120}
	println(numFriendRequests(ages))
}

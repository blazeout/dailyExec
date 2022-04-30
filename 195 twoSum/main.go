package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// hashTable 存储数组中的数字和下标
	hashTable := map[int]int{}
	for i, num := range nums {
		if j, ok := hashTable[target-num]; ok {
			return []int{j, i}
		}
		hashTable[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

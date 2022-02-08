package main

import (
	"fmt"
	"time"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	left, right := 0, 1
	for right < len(nums) {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}

func test() {
	i := 0
ForEnd:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("break now")
				break ForEnd
			}
			fmt.Println("inside the select: ")
		}
		fmt.Println("inside the for: ")
	}
}

func main() {
	test()
}

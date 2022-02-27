package _26_最佳除法

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	// 分子要尽可能的大, 分母要尽可能的小
	var sb strings.Builder
	for i := 0; i < len(nums); i++ {
		if i == 1 {
			sb.WriteString("(")
			sb.WriteString(strconv.Itoa(nums[i]))
			sb.WriteString("/")
		} else if i == len(nums)-1 {
			sb.WriteString(strconv.Itoa(nums[i]))
			sb.WriteString(")")
		} else {
			sb.WriteString(strconv.Itoa(nums[i]))
			sb.WriteString("/")
		}
	}
	return sb.String()
}

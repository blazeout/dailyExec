package _94_最小差值

import (
	"regexp"
	"strconv"
	"strings"
)

func smallestRangeI(nums []int, k int) int {
	maxNum, minNum := nums[0], nums[0]
	// 想要将最大值尽可能小 并且最小值尽可能大
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
		if nums[i] < minNum {
			minNum = nums[i]
		}
	}
	// 如果最大值-k 还要大于 最小值+k, 那么最小差值就是 max-min-2*k
	if maxNum-k > minNum+k {
		return maxNum - minNum - 2*k
	}
	// 其他的情况代表 最大值-k 和 最小值+k 都可以变为一个相等的数
	// 处于这两数之间的数也是可以变为这个相等的数,所以返回0
	return 0
}

func checkOmegaUrl(url string) bool {
	envId := 0
	compile, err := regexp.Compile("env_id=[0-9]+")
	if err != nil {
		return false
	}
	str := compile.FindString(url)
	index := strings.IndexByte(str, '=')
	id, err := strconv.Atoi(str[index+1:])
	if err != nil {
		return false
	}
	envId = id
	if envId == 0 {
		return false
	}
	return true
}

package main

func majorityElement(nums []int) int {
	// 摩尔投票法
	// 遇到一个与当前数相同的数就+1, 不同就-1, 如果count < 0
	count := 1
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			count++
		} else {
			count--
		}
		if count < 0 {
			count = 0
			res = nums[i]
		}
	}
	return res
}

func main() {

}

package main

import "fmt"

func singleNumber(nums []int) int {
	// 统计二进制中的位数, 比如[1,1,1,3], 那么cnt[0] == 3, cnt[1] == 1, 对所有元素进行mod3
	// 最后在 + (1 << i)
	cnt := [32]int{}

	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if num&(1<<i) != 0 {
				cnt[i]++
			}
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		if (cnt[i]%3)&1 == 1 {
			ans += 1 << i
		}
	}
	// golang中int不是32位的, 所以当测试用例出现负数时, 我们如果返回int类型的话, 会将32位看成整数1, 而不是符号位了
	// 所以需要先转成int32类型, 最后再转成int类型的
	return int(int32(ans))
}

func main() {
	ret := singleNumber([]int{2, 2, 2, 3})
	fmt.Println(2 & (1 << 1))
	fmt.Println(ret)
}

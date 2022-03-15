package _49_统计按位或能得到最大值的子集数目

func countMaxOrSubsets(nums []int) int {
	var dfs func(i int, num int)
	max := 0
	ans := -1
	// 每个选项都可以选择选或者不选
	dfs = func(i int, num int) {
		// num := calc(path)
		// if num > max {
		//     max = num
		//     ans = 1
		// }else if num == max{
		//     ans++
		// }
		if i == len(nums) {
			if num > max {
				max = num
				ans = 1
			} else if num == max {
				ans++
			}
			return
		}
		dfs(i+1, num)
		dfs(i+1, num|nums[i])
	}
	dfs(0, 0)
	return ans
}

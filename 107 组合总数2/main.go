package _07_组合总数2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	var dfs func(start int, temp []int, sum int)
	dfs = func(start int, temp []int, sum int) {
		if sum >= target {
			if sum == target {
				t := make([]int, len(temp))
				copy(t, temp)
				res = append(res, t)
			}
			return
		}
		for i := start; i < len(candidates); i++ {
			// 剪枝 同一层已经选择过的数字就不需要了
			if i-1 >= start && candidates[i] == candidates[i-1] {
				continue
			}
			temp = append(temp, candidates[i])
			dfs(i+1, temp, sum+candidates[i])
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0, []int{}, 0)
	return res
}

package _38_组合总数

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	var dfs func(i int, sum int, path []int)
	dfs = func(i int, sum int, path []int) {
		if sum >= target {
			if sum == target {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for j := i; j < len(candidates); j++ {
			path = append(path, candidates[j])
			dfs(j, sum+candidates[j], path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, []int{})
	return res
}

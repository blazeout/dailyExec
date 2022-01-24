package _3_传递信息

func numWays(n int, relation [][]int, k int) (ans int) {
	way := make([][]int, n)
	for _, v := range relation {
		source, dst := v[0], v[1]
		way[source] = append(way[source], dst)
	}
	// ans +1 的条件是count == k, && dfs() == n-1
	var dfs func(count int, index int)
	dfs = func(count int, index int) {
		if count > k {
			return
		}
		if count == k && index == n-1 {
			ans++
		}
		for _, v := range way[index] {
			// 比如0 ,2,4
			dfs(count+1, v)
		}
	}
	dfs(0, 0)
	return ans
}

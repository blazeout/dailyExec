package _40_水壶问题

// 广度优先遍历

// 设计八个状态
// 1. 往A加满水
// 2. 往B加满水
// 3. 把A的水倒完
// 4. 把B的水倒完
// 5.从 A 到 B，使得 B 满，A 还有剩；
// 6.从 A 到 B，此时 A 的水太少，A 倒尽，B 没有满；
// 7. 从 B 到 A，使得 A 满，B 还有剩余；
// 8 .从 B 到 A，此时 B 的水太少，B 倒尽，A 没有满。
var (
	Y int
	Z int
	X int
)

func canMeasureWater(x int, y int, z int) bool {
	// 将x, y, z的值赋给全局变量
	X, Y, Z = x, y, z
	// 定义一个空的map用来保存状态, 因为如果我们回到了重复的状态就没有必要再遍历下去了
	mp := make(map[[2]int]struct{})
	return dfs(&mp, 0, 0)
}

func dfs(mp *map[[2]int]struct{}, xCur, yCur int) bool {
	if xCur == Z || yCur == Z || xCur+yCur == Z {
		return true
	}
	// 这个状态已经出现了 没有必要再遍历下去
	if _, ok := (*mp)[[2]int{xCur, yCur}]; ok {
		return false
	}
	// 记录这个状态
	(*mp)[[2]int{xCur, yCur}] = struct{}{}
	return dfs(mp, X, yCur) || dfs(mp, xCur, Y) ||
		dfs(mp, 0, yCur) || dfs(mp, xCur, 0) ||
		dfs(mp, xCur-min(xCur, Y-yCur), yCur+min(xCur, Y-yCur)) ||
		dfs(mp, xCur+min(yCur, X-xCur), yCur-min(yCur, X-xCur))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

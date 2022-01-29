package _3_地图分析

func maxDistance(grid [][]int) int {
	// 多源BFS,建立虚拟节点, 其实没有必要建立虚拟节点, 只需要将全部的陆地加入队列中即可
	// 方向数组
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := [][]int{}
	m, n := len(grid), len(grid[0])
	// 先将所有的陆地都入队
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	// 从各个陆地开始, 一圈一圈的遍历海洋, 最后遍历到的海洋就是离陆地最远的海洋
	hasOcean := false
	var point []int
	for len(queue) > 0 {
		point = queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		// 取出队列中的第一个元素, 将其四周没有到访过的海洋入队
		for i := 0; i < 4; i++ {
			newX, newY := x+dir[i][0], y+dir[i][1]
			// grid[newX][newY] != 0, 就代表这个海洋被我们访问过
			if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] != 0 {
				continue
			}
			grid[newX][newY] = grid[x][y] + 1
			hasOcean = true
			queue = append(queue, []int{newX, newY})
		}
	}
	// 没有海洋或者没有陆地, 返回-1
	if point == nil || !hasOcean {
		return -1
	}
	return grid[point[0]][point[1]] - 1
}

package _2_地图中的最高点

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	// 我们将陆地==0 水域==1, 反转为陆地==-1, 水域==0
	// 多源BFS, 虚拟节点,将所有的海洋为0的点都入队
	queue := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 0 {
				// 陆地格子不加入队列
				isWater[i][j] = -1

			} else {
				// 海洋格子加入队列
				isWater[i][j] = 0
				queue = append(queue, []int{i, j})
			}
		}
	}
	// 方向数组dir, 向一个点的四个方向扩散
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var point []int
	for len(queue) > 0 {
		point = queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		for i := 0; i < 4; i++ {
			newX, newY := x+dir[i][0], y+dir[i][1]
			// -1代表没有访问过
			if newX < 0 || newX >= m || newY < 0 || newY >= n || isWater[newX][newY] != -1 {
				continue
			}
			isWater[newX][newY] = isWater[x][y] + 1
			queue = append(queue, []int{newX, newY})
		}
	}
	return isWater
}

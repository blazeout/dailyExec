package main

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	// mp := map[[2]int]bool{}
	for i := 0; i < len(dig); i++ {
		x, y := dig[i][0], dig[i][1]
		// 1代表这个地方的工件露出来了
		matrix[x][y] = 1
	}
	ans := 0
next:
	for i := 0; i < len(artifacts); i++ {
		prex, prey := 0, 0
		for j := 0; j < len(artifacts[i]); j += 2 {
			// 对于每个工件内部我们需要测验
			if j == 0 {
				prex, prey = artifacts[i][j], artifacts[i][j+1]
			} else {
				x, y := artifacts[i][j], artifacts[i][j+1]
				if prex == x {
					// 横向扩展的
					for z := prey; z <= y; z++ {
						if matrix[x][z] != 1 {
							continue next
						}
					}
				} else if prey == y {
					// 纵向扩展的
					for z := prex; z <= x; z++ {
						if matrix[z][y] != 1 {
							continue next
						}
					}
				} else {
					if matrix[x-1][y] != 1 || matrix[x][y-1] != 1 || matrix[prex][prey] != 1 || matrix[x][y] != 1 {
						continue next
					}
				}
			}
		}
		ans++
	}
	return ans
}

func main() {
	//5
	//[[3,1,4,1],[1,1,2,2],[1,0,2,0],[4,3,4,4],[0,3,1,4],[2,3,3,4]]
	//[[0,0],[2,1],[2,0],[2,3],[4,3],[2,4],[4,1],[0,2],[4,0],[3,1],[1,2],[1,3],[3,2]]
	digArtifacts(5, [][]int{{3, 1, 4, 1}, {1, 1, 2, 2}, {1, 0, 2, 0}, {4, 3, 4, 4}, {0, 3, 1, 4}, {2, 3, 3, 4}}, [][]int{{0, 0}, {2, 1}, {2, 0}, {2, 3}, {4, 3}, {2, 4}, {4, 1}, {0, 2}, {4, 0}, {3, 1}, {1, 2}, {1, 3}, {3, 2}})

}

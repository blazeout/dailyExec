package main

func countBattleships1(board [][]byte) int {
	m, n := len(board), len(board[0])
	ans := 0
	for i, row := range board {
		for j, _ := range row {
			if board[i][j] == 'X' {
				// 将战舰上下或者左右相邻的X全部设置为'.'
				for k := i + 1; k < m; k++ {
					if board[k][j] == 'X' {
						board[k][j] = '.'
					} else {
						break
					}
				}
				for k := j + 1; k < n; k++ {
					if board[i][k] == 'X' {
						board[i][k] = '.'
					} else {
						break
					}
				}
				ans++
			}
		}
	}
	return ans
}

func countBattleships(board [][]byte) (ans int) {
	// 枚举战舰的左上顶点即可, 因为两艘战舰之间至少有一个水平或垂直的空位分隔
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				// 上面或者左边有板子就不要
				if !(i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X') {
					ans++
				}
			}
		}
	}
	return
}

func main() {

}

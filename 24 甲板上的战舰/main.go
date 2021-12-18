package main

func countBattleships(board [][]byte) int {
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

func main() {

}

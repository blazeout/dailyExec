package _8_机器人能否返回原点

// 字符串模拟即可
func judgeCircle(moves string) bool {
	instance := pair{}
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'L' {
			instance.x++
		} else if moves[i] == 'R' {
			instance.x--
		} else if moves[i] == 'U' {
			instance.y++
		} else {
			instance.y--
		}
	}
	return instance.x == 0 && instance.y == 0
}

type pair struct {
	x int
	y int
}

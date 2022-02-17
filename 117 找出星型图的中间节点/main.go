package _17_找出星型图的中间节点

func findCenter(edges [][]int) int {
	num1, num2 := edges[0][0], edges[0][1]
	num3, num4 := edges[1][0], edges[1][1]
	if num1 == num3 {
		return num1
	} else if num1 == num4 {
		return num1
	}
	if num2 == num3 {
		return num2
	}
	return num2
}

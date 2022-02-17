package _18__字典序的第K小数字_难啊____

func findKthNumber(n int, k int) int {
	//第一个开始的数字
	cur := 1
	// 因为数组下标是从0开始的
	k -= 1
	for k > 0 {
		nodes := getNode(n, cur)
		if k >= nodes {
			k -= nodes
			cur++ // go right
		} else {
			k -= 1
			cur *= 10 // go down
		}
	}
	return cur
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getNode(n int, cur int) int {
	totalNodes := 0
	next := cur + 1
	for cur <= n {
		totalNodes += min(n-cur+1, next-cur)
		next *= 10
		cur *= 10
	}
	return totalNodes
}

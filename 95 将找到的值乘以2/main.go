package _5_将找到的值乘以2

func findFinalValue(nums []int, original int) int {
	mp := map[int]bool{}
	for _, v := range nums {
		mp[v] = true
	}
	isOk := true
	for isOk {
		if mp[original] {
			original *= 2
		} else {
			isOk = false
		}
	}
	return original
}

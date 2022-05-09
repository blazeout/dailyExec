package _97_增减字符串匹配

func diStringMatch(s string) []int {
	res := make([]int, len(s)+1)
	low, high := 0, len(s)
	for i, ch := range s {
		if ch == 'I' {
			res[i] = low
			low++
		} else {
			res[i] = high
			high--
		}
	}
	i := len(s)
	res[i] = low
	return res
}

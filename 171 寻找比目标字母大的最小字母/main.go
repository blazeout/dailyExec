package _71_寻找比目标字母大的最小字母

func nextGreatestLetter(letters []byte, target byte) byte {
	// 二分查找
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	left, right := 0, len(letters)-1
	for left <= right {
		mid := (left + right) / 2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[left]
}

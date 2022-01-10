package _2_峰与谷

func wiggleSort(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		if i&1 == 1 {
			// 此时是谷
			// 判断是否比前面的峰大即可
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else if i&1 == 0 {
			// 此时是峰
			// 判断他比前面的谷是否小即可
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
	return
}

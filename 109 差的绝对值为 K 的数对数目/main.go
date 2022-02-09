package _09_差的绝对值为_K_的数对数目

func countKDifference(nums []int, k int) int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
	}
	count := 0
	for key, value := range mp {
		if p, ok := mp[key+k]; ok {
			count += p * value
		}
	}
	return count
}

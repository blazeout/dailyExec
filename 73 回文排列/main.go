package _3_回文排列

// 只能出现一个奇数个次数的字符, 偶数随意

func canPermutePalindrome(s string) bool {
	mp := map[byte]int{}
	length := len(s)
	for i := 0; i < length; i++ {
		mp[s[i]]++
	}
	once := 0
	for _, value := range mp {
		if value%2 == 1 {
			// 奇数个, 只有中间的一个字符可以
			once++
			if once > 1 {
				return false
			}
		}
	}
	return true
}

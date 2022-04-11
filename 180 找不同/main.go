package _80_找不同

func findTheDifference(s string, t string) byte {
	cnt := [26]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	for j := 0; j < len(t); j++ {
		cnt[t[j]-'a']--
		if cnt[t[j]-'a'] < 0 {
			return t[j]
		}
	}
	return ' '
}

// 神奇求和也行

func findTheDiff(s, t string) byte {
	sum := 0
	for _, v := range s {
		sum -= int(v)
	}
	for _, v := range t {
		sum += int(v)
	}
	return byte(sum)
}

// 异或大法也可以啊 ???
func findTheDifference2(s string, t string) byte {
	var diff byte
	for i := 0; i < len(s); i++ {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}

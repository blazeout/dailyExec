package main

func longestWord(words []string) string {
	mp := map[string]bool{}
	for _, v := range words {
		mp[v] = true
	}
	ans := ""
done:
	for i := 0; i < len(words); i++ {
		for j, _ := range words[i] {
			if _, OK := mp[words[i][:j+1]]; !OK {
				continue done
			}
		}
		// 和答案进行比较
		l1, l2 := len(ans), len(words[i])
		if l1 < l2 || (l1 == l2 && words[i] < ans) {
			ans = words[i]
		}
	}
	return ans
}

func main() {

}

package _79_唯一的摩尔斯密码

var conrespond = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

// a,b,c
func uniqueMorseRepresentations(words []string) int {
	mp := make(map[string]bool)
	count := 0
	for i := 0; i < len(words); i++ {
		tmp := ""
		for j := 0; j < len(words[i]); j++ {
			tmp += conrespond[words[i][j]-'a']
		}
		if _, ok := mp[tmp]; !ok {
			count++
			mp[tmp] = true
		}
	}
	return count
}

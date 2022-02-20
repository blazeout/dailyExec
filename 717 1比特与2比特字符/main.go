package _17_1比特与2比特字符

func isOneBitCharacter(bits []int) bool {
	// 遇到1就走两位, 遇到0就走一位, 判断是否能到达最后一个0即可
	i := 0
	for i < len(bits) {
		if bits[i] == 1 {
			i += 2
			continue
		}
		if i == len(bits)-1 {
			return true
		}
		i++
	}
	return false
}

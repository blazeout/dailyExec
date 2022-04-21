package main

import "unicode"

func mostCommonWord(paragraph string, banned []string) string {
	isBanned := make(map[string]bool)
	for _, b := range banned {
		isBanned[b] = true
	}
	count := map[string]int{}
	var maxCount int
	var word []byte
	for i := 0; i <= len(paragraph); i++ {
		if i < len(paragraph) && unicode.IsLetter(rune(paragraph[i])) {
			word = append(word, byte(unicode.ToLower(rune(paragraph[i]))))
		} else if word != nil {
			str := string(word)
			if !isBanned[str] {
				count[str]++
				maxCount = max(maxCount, count[str])
			}
			word = nil
		}
	}
	for i, v := range count {
		if v == maxCount {
			return i
		}
	}
	return ""
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}

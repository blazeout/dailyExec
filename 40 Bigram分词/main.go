package _0_Bigram分词

import "strings"

func findOcurrences(text string, first string, second string) (res []string) {
	str := strings.Split(text, " ")
	for i := 2; i < len(str); i++ {
		if str[i-2] == first && str[i-1] == second {
			res = append(res, str[i])
		}
	}
	return
}

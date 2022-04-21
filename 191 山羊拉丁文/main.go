package main

import (
	"fmt"
	"strings"
)

func toGoatLatin(sentence string) string {
	if len(sentence) == 0 {
		return ""
	}
	mp := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	const char = "ma"
	index := 1
	sentenceArr := strings.Split(sentence, " ")
	res := ""
	for i := 0; i < len(sentenceArr); i++ {
		temp := sentenceArr[i]
		if _, ok := mp[temp[0]]; ok {
			temp += char + strings.Repeat("a", index)
			res += temp + " "
		} else {
			temp = temp[1:] + string(temp[0]) + char + strings.Repeat("a", index)
			res += temp + " "
		}
		index++
	}
	return res[:len(res)-1]
}

func toGoatLatin2(sentence string) string {
	hash := map[byte]int{
		'A': 1,
		'a': 1,
		'e': 1,
		'E': 1,
		'i': 1,
		'I': 1,
		'o': 1,
		'O': 1,
		'u': 1,
		'U': 1,
	}
	spt := strings.Split(sentence, " ")
	//ret := make([]string, len(spt))
	for i, _ := range spt {
		if hash[spt[i][0]] == 1 {
			sbt := []byte(spt[i])
			sbt = append(sbt, []byte("ma")...)
			for j := 0; j <= i; j++ {
				sbt = append(sbt, "a"...)
			}
			spt[i] = string(sbt)
		} else {
			spt[i] = spt[i][1:] + spt[i][:1]
			sbt := []byte(spt[i])
			sbt = append(sbt, "ma"...)
			for j := 0; j <= i; j++ {
				sbt = append(sbt, "a"...)
			}
			spt[i] = string(sbt)
		}
	}
	return strings.Join(spt, " ")
}

func main() {
	ret := toGoatLatin2("I speak Goat Latin")
	fmt.Println(ret)
}

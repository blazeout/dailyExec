package _9_反转单词前缀

import "strings"

func reversePrefix(word string, ch byte) string {
	index := strings.IndexByte(word, ch)
	if index == -1 {
		return word
	}
	return reverse(word[:index+1]) + word[index+1:]
}

func reverse(s string) string {
	length := len(s)
	a := []byte(s)
	for i := 0; i < length/2; i++ {
		a[i], a[length-1-i] = a[length-1-i], a[i]
	}
	return string(a)
}

package main

import (
	"fmt"
	"strings"
)

func reorderSpaces(text string) string {
	numOfSpace, words := countSpaceAndGetWord(text)
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", numOfSpace)
	}
	remainSpace := numOfSpace % (len(words) - 1)
	averageSpace := numOfSpace / (len(words) - 1)
	var sb strings.Builder
	for _, word := range words {
		sb.WriteString(word)
		sb.WriteString(strings.Repeat(" ", averageSpace))
	}
	return sb.String()[:len(sb.String())-averageSpace] + strings.Repeat(" ", remainSpace)
}

func countSpaceAndGetWord(text string) (count int, words []string) {
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			count++
		}
		if text[i] >= 'a' && text[i] <= 'z' {
			j := i
			for j < len(text) && text[j] != ' ' {
				j++
			}
			words = append(words, text[i:j])
			i = j - 1
		}
	}
	return
}

func main() {
	spaces := reorderSpaces("  this   is  a sentence ")
	fmt.Println(spaces)
}

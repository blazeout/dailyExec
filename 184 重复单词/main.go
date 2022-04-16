package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := ""
	line, _, _ := bufio.NewReader(os.Stdin).ReadLine()
	str = string(line)
	temp := strings.Split(str, " ")
	str = strings.ToLower(str)
	arr := strings.Split(str, " ")

	biaoji := make([]int, len(arr))
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			biaoji[i] = 1
		}
	}
	newStr := ""
	for i := 0; i < len(arr); i++ {
		if biaoji[i] == 0 {
			newStr += temp[i] + " "
		}
	}
	fmt.Println(newStr[:len(newStr)-1])
}

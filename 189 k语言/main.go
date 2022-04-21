package main

import (
	"bufio"
	"fmt"
	"os"
)

func isOk(arr []string) bool {
	//if len(arr)%2 == 1 {
	//	return false
	//}
	var stack []string
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) >= 3 && arr[i][0:3] == "end" && arr[i][3] == ' ' {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last != arr[i][4:] {
				return false
			}
		} else {
			stack = append(stack, arr[i])
		}
	}
	return true
}

func main() {
	m := 0
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		length := 0
		fmt.Scan(&length)
		arr := make([]string, length+1)
		for j := 0; j < length+1; j++ {
			// 按行读取字符串
			read := bufio.NewScanner(os.Stdin)
			read.Scan()
			arr[j] = read.Text()
		}
		if isOk(arr[1:]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

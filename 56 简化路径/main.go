package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	// 栈 , 一个点代表自身, 不需要入栈, 两个点代表要把之前入栈的路径给消灭一个
	// 字母就正常入栈即可
	// "//" 分解之后成为了, 一个空字符串"", 不需要入栈, 因为本来需要变成一个, 后面再join就行
	// return "/" + strings.Join(stack, "/") 就把以/开头成功执行了
	stack := []string{}
	pathArr := split(path, "/")
	for i := 0; i < len(pathArr); i++ {
		if pathArr[i] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if pathArr[i] == "." {
			// 什么也不做
		} else if pathArr[i] == "" {
			// 什么也不做
		} else {
			// 这个就代表是字母了
			stack = append(stack, pathArr[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}

func split(path string, sep string) []string {
	// 以 sep 来分割path字符串
	arr := []string{}
	left, right := 0, 0
	if path[right] == '/' {
		arr = append(arr, "")
	}
	for right < len(path) {
		right++
		for right < len(path) && path[right] != '/' {
			right++
		}
		if right-1 >= left {
			arr = append(arr, path[left+1:right])
		}
		left = right
	}
	return arr
}

func main() {
	path := "/a/./b/../..//c/"
	arr := split(path, "/")
	fmt.Println(arr)
	fmt.Println(len(arr))
}

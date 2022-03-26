package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	var stack []int
	for i := 0; i < len(ops); i++ {
		if ops[i] == "+" {
			m, n := len(stack)-1, len(stack)-2
			stack = append(stack, stack[m]+stack[n])
		} else if ops[i] == "C" {
			// 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数。
			stack = stack[:len(stack)-1]
		} else if ops[i] == "D" {
			// 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
			stack = append(stack, stack[len(stack)-1]*2)
		} else {
			val, _ := strconv.Atoi(ops[i])
			stack = append(stack, val)
		}
	}
	sum := 0
	for i := 0; i < len(stack); i++ {
		sum += stack[i]
	}
	return sum
}

func main() {
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}

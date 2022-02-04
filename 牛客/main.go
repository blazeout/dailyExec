package main

import (
	"fmt"
	"time"
)

// 使用三个channel控制打印流程即可
var ch1 = make(chan bool, 1)
var ch2 = make(chan bool, 1)
var ch3 = make(chan bool, 1)

// 使用两个协程来交替打印数字和字符, 1A2B这种

func f1() {
	// 打印数字
	for i := 1; i <= 26; i++ {
		<-ch1
		time.Sleep(time.Second / 2)
		fmt.Print(i)
		ch2 <- true
	}
}

func f2() {
	for i := 'A'; i <= 'Z'; i++ {
		<-ch2
		fmt.Print(string(i))
		ch1 <- true
	}
	ch3 <- true
}

func main() {
	ch1 <- true
	go f1()
	go f2()
	<-ch3
}

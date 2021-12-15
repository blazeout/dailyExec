package main

import "fmt"

func countPoints(rings string) int {
	b := []byte(rings)
	res := make([]map[byte]bool, 10)
	for i := 0; i < 10; i++ {
		res[i] = make(map[byte]bool)
		res[i]['R'] = false
		res[i]['B'] = false
		res[i]['G'] = false
	}
	for i := 0; i <= len(rings)-2; i += 2 {
		// 为红色的情况下
		res[(b[i+1] - '0')][b[i]] = true
	}
	// res是一个map数组
	count := 0
next:
	for _, mp := range res {
		// value 是一个具体的map
		for _, v := range mp {
			if v == false {
				continue next
			}
		}
		count++
	}
	return count
}

func main() {

	rings := "B0B6G0R6R0R6G9"
	ret := countPoints(rings)
	fmt.Println(ret)
}

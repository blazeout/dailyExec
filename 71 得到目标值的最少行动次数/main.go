package main

import "fmt"

func minMoves(target int, maxDoubles int) int {
	if maxDoubles == 0 {
		return target - 1
	}
	count := 0
	// 贪心, 如果是偶数直接除以2, 如果是奇数那么就减1
	for target > 1 {
		if target%2 == 0 && maxDoubles != 0 {
			// 偶数
			target /= 2
			count++
			maxDoubles--
		} else if target%2 == 1 || maxDoubles == 0 {
			// 奇数
			if maxDoubles != 0 {
				target -= 1
				count++
			} else {
				count += target - 1
				target = 0
			}
		}
	}
	return count
}

func main() {
	ret := minMoves(19, 2)
	fmt.Println(ret)
}

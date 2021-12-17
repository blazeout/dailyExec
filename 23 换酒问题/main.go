package main

import "fmt"

// 自己写的, 思路是全部喝完再换
func numWaterBottles(numBottles int, numExchange int) int {
	count := 0
	emptyBottles := 0
	for numBottles > 0 {
		// 喝了的瓶数
		count += numBottles
		// x代表能换的瓶数
		emptyBottles += numBottles
		x := emptyBottles / numExchange
		// 能换3瓶, 消耗了 num * x
		emptyBottles -= x * numExchange
		numBottles = x
	}
	return count
}

// 官方, 思路是每次bottle的数量大于numExchange就换一瓶, 然后bottle++
func numWaterBottles2(numBottles int, numExchange int) int {
	bottle, ans := numBottles, numBottles
	for bottle >= numExchange {
		bottle -= numExchange
		bottle++
		ans++
	}
	return ans
}

func main() {
	println(numWaterBottles(15, 4))
	fmt.Println(numWaterBottles2(15, 4))
}

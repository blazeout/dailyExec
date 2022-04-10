package main

import (
	"fmt"
	"math"
)

func reverseBits(num uint32) uint32 {
	var sum uint32
	x := 0
	for i := 31; i >= 0; i-- {
		temp := num & (1 << i)
		if temp != 0 {
			sum += 1 * uint32(math.Pow(2.0, float64(x)))
		}
		x++
	}
	return sum
}

func reverseBits2(num uint32) uint32 {
	var cur uint32
	for i := 0; i < 32; i++ {
		cur = (cur << 1) | (num & 1)
		num >>= 1
	}
	return cur
}

func main() {
	fmt.Println(reverseBits(0b00000010100101000001111010011100))
}

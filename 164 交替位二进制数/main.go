package main

func hasAlternatingBits(n int) bool {
	a := n ^ (n >> 1) // 求出全1. 1010 ^ 0101 >== 1111 , 10000 == 0
	return a&(a+1) == 0
}

func main() {

}

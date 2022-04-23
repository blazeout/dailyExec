package main

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	ans := []byte{}
	helper := "0123456789abcdef"
	// 负数右移高位补符号位, 所以ans的长度要小于8
	for num != 0 && len(ans) < 8 {
		temp := num & 15
		ans = append(ans, helper[temp])
		num >>= 4
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}

func main() {
	println(toHex(-1))
}

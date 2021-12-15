package main

import "fmt"

func LongestSubString(str1, str2 string) int {
	// 暴力法, 双重循环固定住一个位置
	m, n := len(str1), len(str2)
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0
			for str1[i+k] == str2[j+k] {
				k++
				if k+i >= m || k+j >= n {
					break
				}
			}
			if k > max {
				max = k
			}
		}
	}
	return max

}

func main() {
	str1 := "abdcefcsa"
	str2 := "abcbefcca"
	ret := LongestSubString(str1, str2)
	fmt.Println("=====")
	fmt.Println(ret)
}

package _10_最简分数

import "fmt"

func simplifiedFractions(n int) []string {
	res := []string{}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if i == 1 {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			} else if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return res
}

// 辗转相除法求最大公约数,
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

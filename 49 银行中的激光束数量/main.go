package _9_银行中的激光束数量

import "strings"

func numberOfBeams(bank []string) int {
	// count1代表上一行出现的激光装置数, count2代表这一行出现的激光装置数
	res := 0 // 代表总的激光束
	count1 := 0
	count2 := 0
	for _, v := range bank[0] {
		if v == '1' {
			count1++
		}
	}
	for i := 1; i < len(bank); i++ {
		if !strings.Contains(bank[i], "1") {
			continue
		}
		for _, v := range bank[i] {
			if v == '1' {
				count2++
			}
		}
		res += count1 * count2
		count1 = count2
		count2 = 0
	}
	return res
}

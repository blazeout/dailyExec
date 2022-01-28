package _0_弱角色的数量

import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0] || properties[i][0] == properties[j][0] && properties[i][1] > properties[j][1]
	})
	length := len(properties)
	res := 0
	maxDefense := properties[length-1][1]
	for i := length - 2; i >= 0; i-- {
		if properties[i][1] >= maxDefense {
			maxDefense = properties[i][1]
		} else {
			res++
		}
	}
	return res
}

package _48_两个列表的最小索引总和

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	mp := map[string]int{}
	res := []string{}
	mp2 := map[string]int{}
	for i := 0; i < len(list1); i++ {
		mp[list1[i]] = i
	}
	for i := 0; i < len(list2); i++ {
		if v, ok := mp[list2[i]]; ok {
			mp2[list2[i]] = i + v
		}
	}
	// key := ""
	value := math.MaxInt32
	for _, v := range mp2 {
		if v < value {
			// key = i
			value = v
		}
	}
	for i, v := range mp2 {
		if v == value {
			res = append(res, i)
		}
	}
	return res
}

// 官解两次循环
func findRestaurant2(list1 []string, list2 []string) []string {
	index := map[string]int{}
	for i := 0; i < len(list1); i++ {
		index[list1[i]] = i
	}
	res := []string{}
	min := math.MaxInt32
	for i, v := range list2 {
		if j, ok := index[v]; ok {
			if i+j < min {
				min = i + j
				res = []string{v}
			} else if i+j == min {
				res = append(res, v)
			}
		}
	}
	return res
}

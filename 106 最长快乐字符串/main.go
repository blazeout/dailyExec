package _06_最长快乐字符串

import "sort"

// 力扣1405题
func longestDiverseString(a int, b int, c int) string {
	ans := []byte{}
	cnt := []struct {
		times int
		char  byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	for {
		// 根据次数来排序
		sort.Slice(cnt, func(i, j int) bool {
			return cnt[i].times > cnt[j].times
		})
		hasNext := false
		for i, v := range cnt {
			if v.times <= 0 {
				break
			}
			m := len(ans)
			if m >= 2 && ans[m-1] == v.char && ans[m-2] == v.char {
				continue
			}
			hasNext = true
			ans = append(ans, v.char)
			cnt[i].times--
			break
		}
		if !hasNext {
			break
		}
	}
	return string(ans)
}

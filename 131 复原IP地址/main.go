package _31_复原IP地址

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	// res返回结果集
	res := []string{}
	var dfs func(start int, subRes []string)
	dfs = func(start int, subRes []string) {
		// 首先将结果加入结果集中, 如果len(subRes) == 4 && start == len(s), 说明全部用上了
		if len(subRes) == 4 && start == len(s) {
			res = append(res, strings.Join(subRes, "."))
			return
		}
		if len(subRes) == 4 && start != len(s) {
			return
		}
		// 一个子串最多截取1~3个长度
		for length := 1; length <= 3; length++ {
			// 如果最后一个子串截取时, 超过了len(s)就需要return
			if start+length-1 >= len(s) {
				return
			}
			// 如果length != 1 && s[start] == '0', 说明包含前导0
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			subRes = append(subRes, str)
			dfs(start+length, subRes)
			subRes = subRes[:len(subRes)-1]
		}
	}
	dfs(0, []string{})
	return res
}

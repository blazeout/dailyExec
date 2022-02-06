package main

import "strconv"

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	// 将所有可行方案全部加入一个数组中
	strArr := []string{}
	// 将余数分为 <= 60 和 > 60的种类
	s := targetSeconds % 60
	m := targetSeconds / 60
	secondM, secondS := s+60, m-1
	strArr = append(strArr, strconv.Itoa(m)+strconv.Itoa(s))
	if secondM >= 0 {
		strArr = append(strArr, strconv.Itoa(secondS)+strconv.Itoa(secondM))
	}
	var check func(s string) bool
	check = func(s string) bool {
		m, _ := strconv.Atoi(s[0:2])
		n, _ := strconv.Atoi(s[2:])
		if (m*60 + n) == targetSeconds {
			return true
		}
		return false
	}
	for i := 0; i < 2; i++ {
		length := len(strArr[i])

		if length == 2 {
			if check("0" + strArr[i]) {
				strArr = append(strArr, "0"+strArr[i])
			}
			if check(strArr[i] + "00") {
				strArr = append(strArr, strArr[i]+"00")
			}
			if check("00" + strArr[i]) {
				strArr = append(strArr, "00"+strArr[i])
			}

			if check(strArr[i] + "0") {
				strArr = append(strArr, strArr[i]+"0")
			}

		} else if length == 3 {
			if check("0" + strArr[i]) {
				strArr = append(strArr, "0"+strArr[i])
			}
			if check(strArr[i] + "0") {
				strArr = append(strArr, strArr[i]+"0")
			}
		}
		if length == 3 {
			if !check("0" + strArr[i]) {
				strArr = append(strArr[:i], strArr[i+1:]...)
			}
		}
	}
	res := 0
	// 下面开始正式的循环计算
	for i := 0; i < len(strArr); i++ {
		cnt := 0
		nowAt := startAt
		if startAt == int(strArr[i][0]-'0') {
			cnt += pushCost
		}
		for j := 1; j < len(strArr[i]); i++ {
			if int(strArr[i][j]-'0') == nowAt {
				cnt += pushCost
			} else {
				cnt += pushCost + moveCost
				nowAt = int(strArr[i][j] - '0')
			}
		}
		res = min(res, cnt)
	}
	return res
}

func min(a, v int) int {
	if a < v {
		return a
	}
	return v
}

func main() {
	minCostSetTime(1, 2, 1, 600)
}

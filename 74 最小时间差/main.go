package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	calc := [][]int{}
	for i := 0; i < len(timePoints); i++ {
		arr := strings.Split(timePoints[i], ":")
		var arr1 []int
		for j := 0; j < len(arr); j++ {
			tmp, _ := strconv.Atoi(arr[j])
			arr1 = append(arr1, tmp)
		}
		calc = append(calc, arr1)
	}
	minRes := math.MaxInt32
	sort.Slice(calc, func(i, j int) bool {
		return calc[i][0]*60+calc[i][1] <= calc[j][0]*60+calc[j][1]
	})
	for i := 1; i < len(calc); i++ {
		tmp := (calc[i][0]*60 + calc[i][1]) - (calc[i-1][0]*60 + calc[i-1][1])
		if tmp == 0 {
			return 0
		}
		minRes = min(minRes, tmp)
	}
	tmp1 := calc[0][0]*60 + 24*60 + calc[0][1]
	tmp2 := calc[len(calc)-1][0]*60 + calc[len(calc)-1][1]
	minRes = min(minRes, abs(tmp1-tmp2))
	minRes = min(minRes, abs(tmp2-tmp1))
	return minRes
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinDifference2(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	cache := make([]int, len(timePoints))
	for i := 0; i < len(timePoints); i++ {
		tmp1, _ := strconv.Atoi(timePoints[i][0:2])
		tmp2, _ := strconv.Atoi(timePoints[i][3:])
		cache[i] = tmp1*60 + tmp2
	}
	sort.Ints(cache)
	minRes := math.MaxInt32
	for i := 1; i < len(cache); i++ {
		minRes = min(minRes, cache[i]-cache[i-1])
	}
	minRes = min(minRes, cache[0]+1440-cache[len(cache)-1])
	return minRes
}

func main() {
	arr := []string{"23:59", "12:01"}
	fmt.Println(arr[0] < arr[1])
	sort.Strings(arr)
	fmt.Println(arr)
	println(findMinDifference(arr))
}

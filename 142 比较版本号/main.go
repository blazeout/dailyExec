package _42_比较版本号

func compareVersion(version1 string, version2 string) int {
	// 我们每次将'.'号, 前面的转换成数字, 然后比较即可
	m, n := len(version1), len(version2)
	i, j := 0, 0
	for i < m || j < n {
		var x, y int
		// x代表version1的局部数字, y代表version2的局部数字
		for ; i < m && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		for ; j < n && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		// 跳过两个点
		i, j = i+1, j+1
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

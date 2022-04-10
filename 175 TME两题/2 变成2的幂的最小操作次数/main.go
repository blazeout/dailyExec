package __变成2的幂的最小操作次数

// 思路 : 2的幂肯定是二进制位只有一个1才是2的幂
// 所以我们统计字符串中的1的个数即可, 然后return count - 1
// 比如 "1011" -> 变成 "0001" 就是2的0次方幂
func count(str string) int {
	countOne := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			countOne++
		}
	}
	return countOne - 1
}

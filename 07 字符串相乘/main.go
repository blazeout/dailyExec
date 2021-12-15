package _7_字符串相乘

import "strconv"

// func multiply(num1 string, num2 string) string {
//     if num1 == "0" || num2 == "0" {
//         return "0"
//     }
//     ans := "0"
//     m, n := len(num1), len(num2)
//     // 求出被乘数和乘数的每一位相乘的结果 比如 456 乘以 123
//     // 就是 首先"0" + 456 * 3 + 456 * 20 + 100 * 456
//     for i := n-1; i >= 0; i-- {
//         curr := "" // 暂时储存目前这一位的数与被乘数的乘积
//         carry := 0 // 用来记录进位
//         for j := n-1; j > i; j-- {
//             curr += "0" // 第一位就不需要加0, 当到第二位时,就需要加一位 20
//         }
//         y := int(num2[i]-'0')
//         for j := m-1; j >= 0; j-- {
//             x := int(num1[j]-'0')
//             result := x*y+carry
//             curr = strconv.Itoa(result%10)+curr
//             carry = result/10
//         }
//         // 将这里处理完后, 可能carry进位不为0, 需要将这一位加入curr中
//         for ; carry != 0; carry /= 10 {
//             curr = strconv.Itoa(carry%10) + curr
//         }
//         ans = addString(ans, curr)
//     }
//     return ans
// }

// func addString(num1 string, num2 string) string {
//     ans := ""
//     carry := 0 // 进位
//     for m, n := len(num1)-1, len(num2)-1; m >= 0 || n >= 0 || carry != 0; m, n = m-1, n-1 {
//         var x,y int // 当m或者n中有一个位数不够是, 就默认为0
//         if m >= 0 {
//             x = int(num1[m]-'0')
//         }
//         if n >= 0 {
//             y = int(num2[n]-'0')
//         }
//         result := x+y+carry
//         ans = strconv.Itoa(result%10)+ans
//         carry = result/10
//     }
//     return ans
// }


// 方法1采用了太多的字符串相加和相乘, 如果我们将结果存于数组中, 那么速度就会快很多

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := ""
	// m, n分别为num1, num2的长度, 则总的字符串长度为m+n-1或者m+n
	m, n := len(num1), len(num2)
	arr := make([]int, m+n)
	for i := m-1; i >= 0; i-- {
		x := int(num1[i]-'0')
		for j := n-1; j >= 0; j-- {
			y := int(num2[j]-'0')
			arr[i+j+1] += x*y
		}
	}
	// 将进位处理
	for i := m+n-1; i > 0; i-- {
		arr[i-1] += arr[i]/10
		arr[i] %= 10
	}
	index := 0
	if arr[index] == 0 {
		index++
	}
	for ; index < m+n; index++ {
		ans += strconv.Itoa(arr[index])
	}
	return ans
}


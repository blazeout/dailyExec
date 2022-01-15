package _9_计算力扣银行的钱

// 简单模拟计算
func totalMoney(n int) int {
	sum := 0
	count1 := 1
	newCount := count1
	for i := 0; i < n; i++ {
		if i%7 == 0 {
			count1 = newCount
			newCount = count1 + 1
		}
		sum += count1
		count1++
	}
	return sum
}

// 等差数列计算

func totalMoney2(n int) (ans int) {
	// 所有完整的周存的钱
	weekNum := n / 7
	firstWeekMoney := (1 + 7) * 7 / 2
	lastWeekMoney := firstWeekMoney + 7*(weekNum-1)
	weekMoney := (firstWeekMoney + lastWeekMoney) * weekNum / 2
	// 剩下的不能构成一个完整的周的天数里存的钱
	dayNum := n % 7
	firstDayMoney := 1 + weekNum
	lastDayMoney := firstDayMoney + dayNum - 1
	dayMoney := (firstDayMoney + lastDayMoney) * dayNum / 2
	return weekMoney + dayMoney
}

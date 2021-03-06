package _54_Nim游戏

func canWinNim(n int) bool {
	// 脑筋急转弯
	// 结论是谁移走了最后一块石头就是谁赢了
	// 如果是1~3 先手必赢
	// 如果是4, 先手必输
	// 如果是5~7, 我们可以将石头的数量变为4, 这样就相当于后手变成了4个的先手 他必输
	// 如果是8, 不论我们怎么取, 对手都可以将石头变为4个, 我们必输
	// 所以得出结论, 如果是4的倍数, 那么就是我们必输
	return n%4 != 0
}

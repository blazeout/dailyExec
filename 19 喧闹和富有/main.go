package _9_喧闹和富有

// 好题!!!
func loudAndRich(richer [][]int, quiet []int) []int {
	// 拓扑排序, 建图, 建立入度, 根据将入度为0的先放入队列
	T := make([][]int, len(quiet))      // 领接表 图
	inDegree := make([]int, len(quiet)) // 入度表
	// 将钱多的指向钱少的, 在处理钱多的时候, 钱少的quiet值如果小于钱多的quiet值, 那么就可以更新钱少的ans值
	for _, v := range richer {
		T[v[0]] = append(T[v[0]], v[1]) // 这一步将钱少的人加入到钱多的人里面的领接表中
		inDegree[v[1]]++                // 这一步将钱少的人的读入加1
	}
	// 进过上述的步骤就已经将领接表和入度表构建好了
	ans := make([]int, len(quiet)) // 在拓扑排序中, 逐一更新钱少的人的ans
	for i := range ans {
		ans[i] = i // 初始化, 入度为0的人的ans肯定是本身, 因为没有比他更有钱的人了
	}
	// 使用队列进行拓扑排序
	q := make([]int, 0)
	// 将钱多的人加入队列中, 即入度为0的人
	for i, v := range inDegree {
		if v == 0 {
			q = append(q, i)
		}
	}
	// 下面就是进行拓扑排序
	for len(q) > 0 {
		cur := q[0] // 取出第一人
		q = q[1:]
		for _, v := range T[cur] {
			// v 代表钱比cur 少的人, 比如 领接表T中的3, 那么他就是这样的 T[4] = [1, 7] v就是其中的1,7
			// 在此过程中更新钱少的人即v人的ans, 因为cur是比他钱多的人, 如果quiet[cur] < quite[v]
			// 那么就将ans[v] = ans[cur], 上面有一个坑, 就是quiet[cur] < quite[v]是错的
			// 因为如果是这样比就只是比较这两个人quiet了, 我们需要找出最安静的人, 所以要传入
			// quiet[ans[cur]] < quiet[ans[v]] ans代表着第i个上最安静的人, 所以需要这样比较
			if quiet[ans[cur]] < quiet[ans[v]] {
				ans[v] = ans[cur]
			}
			// 此时将v的入度减1, 代表对于cur, v已经遍历过了
			inDegree[v]--
			// 当in_degree[v] == 0 , 代表他此时就是钱多的人了 加入队列
			if inDegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return ans
}

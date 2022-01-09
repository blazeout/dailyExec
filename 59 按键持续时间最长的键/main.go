package _9_按键持续时间最长的键

func slowestKey(releaseTimes []int, keysPressed string) byte {
	answer := keysPressed[0]
	maxTime := releaseTimes[0]
	for i := 1; i < len(keysPressed); i++ {
		key := keysPressed[i]
		duringTime := releaseTimes[i] - releaseTimes[i-1]
		if duringTime > maxTime || duringTime == maxTime && key > answer {
			maxTime = duringTime
			answer = key
		}
	}
	return answer
}

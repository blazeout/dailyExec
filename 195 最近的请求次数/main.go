package _95_最近的请求次数

const interval = 3000

// RecentCounter 使用数组和length来求解
type RecentCounter struct {
	nums   []int
	length int
}

func Constructor() RecentCounter {
	return RecentCounter{
		nums:   make([]int, 0),
		length: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.nums = append(this.nums, t)
	this.length++
	left := t - interval
	if left < 0 {
		left = 0
	}
	// 二分查找
	index := BinarySearch(this.nums, left)
	return this.length - index
}

func BinarySearch(nums []int, target int) int {
	// 返回target的下标
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

type RecentCounter2 struct {
	queue []int
}

func Constructor2() RecentCounter2 {
	return RecentCounter2{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter2) Ping2(t int) int {
	this.queue = append(this.queue, t)
	for this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}
	return len(this.queue)
}

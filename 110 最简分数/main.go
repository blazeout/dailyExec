package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

type Instance struct {
}

var instance *Instance
var lock sync.Mutex
var lock1 sync.RWMutex
var x sync.Map

func GetInstance() *Instance {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = new(Instance)
		}
	}
	return instance
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := Patition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func Patition(nums []int, left, right int) int {
	random := rand.Intn(right-left+1) + left
	nums[right], nums[random] = nums[random], nums[right]
	pivot := nums[right]
	l := left - 1
	r := left
	for ; r < right; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l+1], nums[right] = nums[right], nums[l+1]
	return l + 1
}

func simplifiedFractions(n int) []string {
	res := []string{}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if i == 1 {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			} else if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", i, j))
			}
		}
	}
	return res
}

// 辗转相除法求最大公约数,
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	nums := []int{9, 4, 1, 7}
	quickSort(nums, 0, 3)
	fmt.Println(nums)
}

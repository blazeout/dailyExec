package main

type NumArray struct {
	arr    []int // 本身的数组
	prefix []int // 前缀数组
}

func Constructor(nums []int) NumArray {
	prefixArr := make([]int, len(nums)+1)
	sum := nums[0]
	for i := 1; i < len(prefixArr); i++ {
		prefixArr[i] = sum
		if i == len(nums) {
			break
		}
		sum += nums[i]
	}
	return NumArray{
		arr:    nums,
		prefix: prefixArr,
	}
}

func (a *NumArray) Update(index int, val int) {
	before := a.arr[index]
	a.arr[index] = val
	for i := index + 1; i < len(a.prefix); i++ {
		a.prefix[i] += val - before
	}
}

// SumRange 前缀数组 ?
func (a *NumArray) SumRange(left int, right int) int {
	return a.prefix[right+1] - a.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

func main() {
	numArray := Constructor([]int{1, 3, 5})
	numArray.SumRange(0, 2) // 返回 1 + 3 + 5 = 9
	numArray.Update(1, 2)   // nums = [1,2,5]
	numArray.SumRange(0, 2) // 返回 1 + 2 + 5 = 8
}

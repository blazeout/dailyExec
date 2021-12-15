package _2_求众数

// 1. 最直观解法 求出所有元素出现的次数判断是否有大于n/3的元素
// 时间O(n), 空间O(n)
func majorityElement(nums []int) []int {
    m1 := make(map[int]int)
    n := len(nums)/3.0
    for  _ , num := range nums {
        m1[num]++
    }
    res := []int{}
    for key , value := range m1 {
        if value > n {
            res = append(res, key)
        }
    }
    return res
}

// 2.摩尔投票法
func majorityElement2(nums []int)(ans []int) {
    element1, element2 := 0, 0
    vote1, vote2 := 0,0
    for _, num := range nums {
        if vote1 >0 && num == element1 {
            vote1++ // 当前元素为众数1就将vote1++
        }else if vote2 > 0 && num == element2 {
            vote2++ // 当前元素为众数2就将vote2++
        }else if vote1 == 0 {
            element1 = num // 等于0就将当前元素设置为众数1
            vote1++
        }else if vote2 == 0 { // vote2等于0将当前元素设置为众数2
            element2 = num
            vote2++
        }else {
            vote1--
            vote2--
        }
    }
    cnt1, cnt2 := 0, 0
    for  _ , v  := range nums {
        if vote1 > 0 && v == element1 {
            cnt1++
        }else if vote2 > 0 && v == element2{
            cnt2++
        }
    }
    if vote1 > 0 && cnt1 > len(nums)/3 {
        ans = append(ans, element1)
    }
    if vote2 > 0 && cnt2 > len(nums)/3 {
        ans = append(ans, element2)
    }
    return
}

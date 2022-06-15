package main

import "fmt"

// 删列造序
func minDeletionSize(strs []string) int {
	count := 0
next:
	for j := 0; j < len(strs[0]); j++ {
		before := strs[0][j]
		for i := 1; i < len(strs); i++ {
			if before > strs[i][j] {
				count++
				continue next
			} else {
				before = strs[i][j]
			}
		}
	}
	return count
}

type Student struct {
	Name string
	Age  int
}

func main() {
	arr := []int{1, 2, 3}
	arr = append(arr, 4, 5, 6, 7)
	fmt.Println(cap(arr), len(arr))
	s1 := Student{"a", 1}
	s2 := Student{"b", 2}
	s3 := Student{"c", 3}
	arr2 := []Student{s1}
	arr2 = append(arr2, s2, s3)
	fmt.Println(cap(arr2))
}

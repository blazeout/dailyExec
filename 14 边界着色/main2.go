package main

import (
	"fmt"
	"unsafe"
)



func main() {
	grid := make([][]int, 3)
	for i := range grid {
		grid[i] = make([]int, 3)
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = 1
		}
	}
	revise := make([][]int, 3)
	for i := range revise {
		revise[i] = make([]int, 3)
	}
	copy(revise, grid)
	fmt.Println(unsafe.Pointer(&grid))
	fmt.Println(unsafe.Pointer(&revise))
}

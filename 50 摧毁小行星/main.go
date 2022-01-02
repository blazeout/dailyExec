package _0_摧毁小行星

// 贪心 每次都碰撞最小的行星
func asteroidsDestroyed(mass int, asteroids []int) bool {
	heapSort(asteroids)
	for i := 0; i < len(asteroids); i++ {
		if mass >= asteroids[i] {
			mass += asteroids[i]
		} else {
			return false
		}
	}
	return true
}

// 堆排序

func heapSort(arr []int) {
	size := len(arr)
	createHeap(arr, size)
	for i := size - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		size--
		shiftDown(arr, 0, size)
	}
}

func createHeap(arr []int, size int) {
	root := (len(arr) - 2) >> 1
	for ; root >= 0; root-- {
		shiftDown(arr, root, size)
	}
}

func shiftDown(arr []int, parent int, size int) {
	child := parent*2 + 1
	for child < size {
		if child+1 < size && arr[child] < arr[child+1] {
			child += 1
		}
		if arr[parent] < arr[child] {
			arr[parent], arr[child] = arr[child], arr[parent]
		} else {
			break
		}
		parent = child
		child = 2*parent + 1
	}
}

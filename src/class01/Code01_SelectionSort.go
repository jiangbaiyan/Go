package class01

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 1. 这里必须记录minIndex, 不能直接记录值, 否则不会达到交换的效果
	// 2. 优化: 可以在外层循环 < len(arr) - 1. 因为倒数第二个数有序了, 最后一个数也一定有序了
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

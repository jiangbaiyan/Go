package class01

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	if arr == nil || len(arr) == 0 {
		return
	}
	// i控制0~0有序 0~1有序 0~2...0~i有序
	for i := 1; i < len(arr); i++ {
		// j永远在i的位置, 和左边那个数做比较, 如果小就交换, 直到比他大或者没有数再停止
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

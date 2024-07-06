package class01

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	// 每次冒泡都会把最大的排到最后e的位置上
	for e := len(arr) - 1; e > 0; e-- {
		for i := 0; i < e; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}

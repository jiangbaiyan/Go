package class08

// CountSort 计数排序, 适用于如排序员工年龄(最大不超过200)
func CountSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 这里简化, 假设最大值为200
	bucket := make([]int, 200, 200)
	for _, item := range arr {
		bucket[item]++
	}
	// 原数组下标
	i := 0
	for num, item := range bucket {
		for j := 0; j < item; i++ {
			arr[i] = num
			i++
		}
	}
}

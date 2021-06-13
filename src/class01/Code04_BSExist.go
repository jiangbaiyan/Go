package class01

func BSExist(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	L := 0
	R := len(arr) - 1
	mid := 0
	// 如果这里包含了等于,那么可以直接return false
	for L <= R {
		mid = L + ((R - L) >> 1)
		if num > arr[mid] {
			L = mid + 1
		} else if num < arr[mid] {
			R = mid - 1
		} else {
			return true
		}
	}
	return false
}

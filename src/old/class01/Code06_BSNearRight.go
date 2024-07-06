package class01

// BSNearRight 在arr上，找满足<=value的最右位置
func BSNearRight(arr []int, value int) int {
	L := 0
	R := len(arr) - 1
	mid := 0
	index := -1
	for L <= R {
		mid = L + (R-L)>>1
		if arr[mid] <= value {
			index = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return index
}

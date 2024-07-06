package class01

// BSAwesome 局部最小值问题
// 无序数组 任意相邻的数字不相等
func BSAwesome(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	// 0~1是上扬的，那么0就是局部最小值
	if len(arr) == 0 || arr[0] < arr[1] {
		return 0
	}
	// 倒数第二个~倒数第一个下降的，那么倒数第一个就是局部最小值
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}
	// 0~1下降，倒数第二~倒数第一个上扬，那么中间必然存在一个局部最小值点
	L := 1            // 注意，L从1开始
	R := len(arr) - 2 // 注意，R从len - 2开始
	mid := 0
	// 注意，这里不能用L <= R，因为mid - 1可能会越界
	for L < R {
		mid = L + (R-L)>>1
		if arr[mid] > arr[mid-1] {
			// 中点比左边那个数大，局部最小在左边
			R = mid - 1
		} else if arr[mid] > arr[mid+1] {
			// 中点比右边那个数大，局部最小在左边
			L = mid + 1
		} else {
			return mid
		}
	}
	return arr[L]
}

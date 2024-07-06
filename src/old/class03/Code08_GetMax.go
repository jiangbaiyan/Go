package class03

// GetMax 获取数组中最大值，递归实现
func GetMax(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

// 在L..R上求arr的最大值
func process(arr []int, L int, R int) int {
	// base case
	if L == R {
		return arr[L]
	}
	// 分解子问题, 左边求最大值, 右边求最大值, 然后二者求max就是整体的最大值
	mid := L + ((R - L) >> 1)
	return max(process(arr, L, mid), process(arr, mid+1, R))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

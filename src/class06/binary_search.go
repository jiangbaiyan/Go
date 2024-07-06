package class06

// 有序数组二分法
func findNumber(arr []int, num int) bool {
	if arr == nil || len(arr) == 0 {
		return false
	}
	l := 0
	r := len(arr)
	for l <= r {
		m := l + (r-l)>>1
		if arr[m] == num {
			return true
		} else if arr[m] < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false
}

// 有序数组找 >=num 的最左下标
func findLeft(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	ans := -1
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := l + (r-l)>>1
		if arr[m] >= num {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}

// 有序数组找 <=num 的最右下标
func findRight(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	ans := -1
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := l + (r-l)>>1
		if arr[m] <= num {
			ans = m
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}

// https://leetcode.cn/problems/find-peak-element/description/
// 峰值元素是指其值严格大于左右相邻值的元素。
//
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
// 你可以假设 nums[-1] = nums[n] = -∞ 。
//
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
func findPeakElement(arr []int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return 0
	}
	n := len(arr)
	if arr[0] > arr[1] {
		return 0
	}
	if arr[n-1] > arr[n-2] {
		return n - 1
	}

	// 0不是峰值，n-1不是峰值，那么1...n-2必有峰值（因为-1-0向上，n-1-n向下）
	l := 1
	r := n - 2
	for l <= r {
		m := l + (r-l)>>1
		if arr[m] < arr[m+1] {
			l = m + 1
		} else if arr[m-1] > arr[m] {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}

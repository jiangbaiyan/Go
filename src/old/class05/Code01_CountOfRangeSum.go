package class05

// CountOfRangeSum
// https://leetcode.com/problems/count-of-range-sum/
//
// 给定一个数组arr，两个整数lower和upper，
//
// 返回arr中有多少个子数组的累加和在[lower,upper]范围上
func CountOfRangeSum(arr []int, lower int, upper int) int {
	if len(arr) == 0 {
		return 0
	}
	sum := make([]int, len(arr), len(arr))
	sum[0] = arr[0]
	for i, item := range arr {
		sum[i] = sum[i-1] + item
	}
	return process(sum, 0, len(sum)-1, lower, upper)
}

func process(sum []int, L int, R int, lower int, upper int) int {
	if L == R {
		if sum[L] >= lower && sum[L] <= upper {
			return 1
		} else {
			return 0
		}
	}
	mid := L + ((R - L) >> 1)
	return process(sum, L, mid, lower, upper) + process(sum, mid+1, R, lower, upper) + merge(sum, L, mid, R, lower, upper)
}

func merge(sum []int, L int, mid int, R int, lower int, upper int) int {
	ans := 0
	windowL := L
	windowR := L
	for i := mid + 1; i <= R; i++ {
		min := sum[i] - upper
		max := sum[i] - lower
		for windowL <= mid && windowL < min {
			windowL++
		}
		for windowR <= mid && windowR <= max {
			windowR++
		}
		ans += windowR - windowL
	}

	help := make([]int, R-L+1, R-L+1)
	p1 := L
	p2 := mid + 1
	i := 0
	for p1 <= mid && p2 <= R {
		if sum[p1] <= sum[p2] {
			help[i] = sum[p1]
			p1++
			i++
		} else {
			help[i] = sum[p2]
			p2++
			i++
		}
	}
	for p1 <= mid {
		help[i] = sum[p1]
		p1++
		i++
	}
	for p2 <= R {
		help[i] = sum[p2]
		p2++
		i++
	}
	for i, item := range help {
		sum[L+i] = item
	}
	return ans
}

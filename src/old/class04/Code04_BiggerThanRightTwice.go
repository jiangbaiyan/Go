package class04

// BiggerThanRightTwice 在一个数组中，
// 对于每个数num，求有多少个后面的数 * 2 依然<num，求总个数
// 比如：[3,1,7,0,2]
// 3的后面有：1，0
// 1的后面有：0
// 7的后面有：0，2
// 0的后面没有
// 2的后面没有
// 所以总共有5个
func BiggerThanRightTwice(arr []int) int {
	return process4(arr, 0, len(arr)-1)
}

func process4(arr []int, L int, R int) int {
	if L == R {
		return 0
	}
	mid := L + ((R - L) >> 1)
	return process4(arr, L, mid) + process4(arr, mid+1, R) + merge4(arr, L, mid, R)
}

func merge4(arr []int, L int, mid int, R int) int {

	ans := 0
	windowR := mid + 1
	for i := L; i <= mid; i++ {
		for windowR <= R && (arr[i] > 2*arr[windowR]) {
			windowR++
		}
		ans += windowR - mid - 1
	}

	help := make([]int, R-L+1, R-L+1)
	p1 := L
	p2 := mid + 1
	i := 0

	for p1 <= mid && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			i++
			p1++
		} else {
			help[i] = arr[p2]
			i++
			p2++
		}
	}

	for p1 <= mid {
		help[i] = arr[p1]
		p1++
		i++
	}

	for p2 <= R {
		help[i] = arr[p2]
		p2++
		i++
	}

	for i, item := range help {
		arr[L+i] = item
	}

	return ans
}

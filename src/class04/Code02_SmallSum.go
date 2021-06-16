package class04

// SmallSum 在一个数组中，一个数左边比它小的数的总和，叫数的小和，所有数的小和累加起来，叫数组小和。求数组小和。
// 例子： [1,3,4,2,5]
// 1左边比1小的数：没有
// 3左边比3小的数：1
// 4左边比4小的数：1、3
// 2左边比2小的数：1
// 5左边比5小的数：1、3、4、 2
// 所以数组的小和为1+1+3+1+1+3+4+2=16
func SmallSum(arr []int) int {
	return process2(arr, 0, len(arr)-1)
}

// 在L..R上求小和, 并将结果返回
func process2(arr []int, L int, R int) int {
	if L == R {
		return 0
	}
	mid := L + ((R - L) >> 1)
	return process2(arr, L, mid) + process2(arr, mid+1, R) + merge2(arr, L, mid, R)
}

func merge2(arr []int, L int, mid int, R int) int {
	help := make([]int, R-L+1, R-L+1)
	p1 := L
	p2 := mid + 1
	i := 0
	res := 0
	for p1 <= mid && p2 <= R {
		// p1 < p2 拷贝左组, 等于拷贝右组
		if arr[p1] < arr[p2] {
			help[i] = arr[p1]
			// 记录小和
			res += arr[p1] * (R - p2 + 1)
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
		i++
		p1++
	}
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}
	for i, item := range help {
		arr[L+i] = item
	}
	return res
}

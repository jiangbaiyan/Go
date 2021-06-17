package class04

func ReversePair(arr []int) int {
	return process3(arr, 0, len(arr)-1)
}

func process3(arr []int, L int, R int) int {
	if L == R {
		return 0
	}
	mid := L + (R-L)>>1
	return process3(arr, 0, mid) + process3(arr, mid+1, R) + merge3(arr, L, mid, R)
}

func merge3(arr []int, L int, mid int, R int) int {
	help := make([]int, R-L+1, R-L+1)
	p1 := mid
	p2 := R
	i := len(help) - 1
	res := 0
	for p1 >= L && p2 >= mid+1 {
		// 相等先拷贝右组
		if arr[p2] >= arr[p1] {
			help[i] = arr[p2]
			i--
			p2--
		} else {
			// 拷贝左组 产生逆序对
			res = res + (p2 - mid)
			help[i] = arr[p1]
			i--
			p1--
		}
	}
	for p1 >= L {
		help[i] = arr[p1]
		i--
		p1--
	}
	for p2 >= mid+1 {
		help[i] = arr[p2]
		i--
		p2--
	}
	for i, item := range help {
		arr[L+i] = item
	}
	return res
}

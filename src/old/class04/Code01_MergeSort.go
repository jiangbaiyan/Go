package class04

func MergeSort(arr []int) {
	process1(arr, 0, len(arr)-1)
}

// 子过程, 数组在L..R上有序
// 分解子问题: 砍一半, 左边和右边分别有序, 然后merge就是整体有序
func process1(arr []int, L int, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)>>1
	process1(arr, L, mid)   // 让左边有序
	process1(arr, mid+1, R) // 让右边有序
	// 让整体有序
	merge(arr, L, mid, R)
}

// 左边有序 右边有序 让L..R范围在整体上有序的merge操作
func merge(arr []int, L int, mid int, R int) {
	// 准备一个辅助数组
	help := make([]int, R-L+1, R-L+1)
	// 两个指针, 一个指向L, 一个指向mid + 1
	p1 := L       // 左指针
	p2 := mid + 1 // 右指针
	i := 0        // 辅助数组下标
	// 两个指针都没遍历完成
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
	// p1指针未遍历完成
	for p1 <= mid {
		help[i] = arr[p1]
		i++
		p1++
	}
	// p2指针未遍历完成
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}
	// 写回原数组
	for i, item := range help {
		arr[L+i] = item // 注意这里必须是L + i, 不能直接写i, 因为help的下标永远从0开始, 而原数组有可能不是从0开始(因某个子过程下标并非从0开始)
	}
	return
}

// MergeSort2 非递归版归并排序
func MergeSort2(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	N := len(arr)
	mergeSize := 1
	for mergeSize < N {
		// 当前左组的第一个位置
		L := 0
		for L < N {
			M := L + mergeSize - 1
			// 没有右组, 直接不管
			if M >= N {
				break
			}
			// 有右组, 右组也可能越界, 取不越界的那个
			R := min(M+mergeSize, N-1)
			merge(arr, L, M, R)
			// 继续下一组
			L = R + 1
		}
		// 防止溢出
		if mergeSize > N/2 {
			break
		}
		// 步长 * 2
		mergeSize <<= 1
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

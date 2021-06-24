package class06

// SortArrayDistanceLessK 已知一个几乎有序的数组。几乎有序是指，如果把数组排好顺序的话，每个元素移动的距离一定不超过k，并且k相对于数组长度来说是比较小的。
//请选择一个合适的排序策略，对这个数组进行排序。
// 例: [3, 4, 1, 2, 5] 排好序后为 [1, 2, 3, 4, 5]
func SortArrayDistanceLessK(arr []int, k int) {
	if k == 0 {
		return
	}
	// 这里要用最小堆, 没有实现所以先这样代替
	h := &MyHeap{}
	index := 0
	// index走过0...k - 1, 并构建起一个堆
	for ; index < min(len(arr)-1, k-1); index++ {
		h.Push(arr[index])
	}
	// index来到k位置上, 标识堆的结束位置
	// i从0开始, 标识要赋值给原数组的位置
	i := 0
	for ; index < len(arr); index++ {
		h.Push(arr[index])
		arr[i] = h.Pop()
		i++
	}
	for !h.IsEmpty() {
		arr[i] = h.Pop()
		i++
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

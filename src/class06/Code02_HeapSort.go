package class06

func HeapSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 先构建最大堆

	/*	第一种方式, 来一个数heapInsert一次, 复杂度 N * logN
		for i := range arr {
			HeapInsert(arr, i)
		}
	*/

	// 第二种方式, 从最后一个往前heapify, 复杂度O(N)
	for i := len(arr) - 1; i >= 0; i-- {
		Heapify(arr, i, len(arr))
	}
	// 堆建好了, 那么第一个位置就是最大值, 需要和最后一个位置交换
	heapSize := len(arr)
	arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
	// 最后一个数排好序了, 堆缩小一个
	heapSize--
	// 此时堆顶并不是最大的元素, 需要heapify下
	for heapSize > 0 {
		Heapify(arr, 0, heapSize)
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
		heapSize--
	}
}

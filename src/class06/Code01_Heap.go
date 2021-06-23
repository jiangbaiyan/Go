package class06

type MyHeap struct {
	heap     []int
	heapSize int
}

// HeapInsert i位置上的数加进来，向上串，直到堆调整好
func (h *MyHeap) HeapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

// Heapify index位置的数向下看
func (h *MyHeap) Heapify(arr []int, index int, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		// 左右比大小
		var maxChild int
		if left+1 < heapSize && arr[left+1] > arr[left] {
			maxChild = left + 1
		} else {
			maxChild = left
		}
		if arr[index] > arr[maxChild] {
			arr[index], arr[maxChild] = arr[maxChild], arr[index]
			index = maxChild
			left = index*2 + 1
		}
	}
}

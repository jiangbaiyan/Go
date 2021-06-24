package class06

type MyHeap struct {
	heap     []int
	heapSize int
}

// HeapInsert i位置上的数加进来，向上串，直到堆调整好
func HeapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

// Heapify index位置的数向下看
func Heapify(arr []int, index int, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		// 左右比大小
		var maxChild int
		if left+1 < heapSize && arr[left+1] > arr[left] {
			maxChild = left + 1
		} else {
			maxChild = left
		}
		// 如果当前值比孩子的最大值还大, 说明它就是最大值了, 退出循环
		if arr[index] > arr[maxChild] {
			break
		}
		// 如果孩子更大, 那么交换
		arr[index], arr[maxChild] = arr[maxChild], arr[index]
		// 当前值来到child的位置, 继续下一轮循环
		index = maxChild
		left = index*2 + 1
	}
}

func (h *MyHeap) Push(value int) {
	h.heap = append(h.heap, value)
	h.heapSize++
	HeapInsert(h.heap, h.heapSize)
}

func (h *MyHeap) Pop() int {
	ans := h.heap[0]
	h.heap[0], h.heap[h.heapSize-1] = h.heap[h.heapSize-1], h.heap[0]
	Heapify(h.heap, 0, h.heapSize)
	return ans
}

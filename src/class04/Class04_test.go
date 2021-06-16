package class04

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{3, 2, 4, 5, 1, 9, 4}
	MergeSort(arr)
	t.Log(arr)
	arr = []int{3, 2, 4, 5, 1, 9, 4}
	MergeSort2(arr)
	t.Log(arr)
}

func TestSmallSum(t *testing.T) {
	arr := []int{1, 3, 4, 2, 5}
	t.Log(SmallSum(arr))
}

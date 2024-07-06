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

func TestReversePair(t *testing.T) {
	arr := []int{2, 5, 1, 4}
	t.Log(ReversePair(arr))
}

func TestBiggerThanRightTwice(t *testing.T) {
	arr := []int{3, 1, 7, 0, 2}
	t.Log(BiggerThanRightTwice(arr))
}

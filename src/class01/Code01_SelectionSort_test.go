package class01

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := []int{6, 3, 4, 2, 1}
	SelectionSort(arr)
	t.Log(arr)
}

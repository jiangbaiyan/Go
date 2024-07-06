package class06

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	// 测试findNumber函数
	arr1 := []int{1, 3, 5, 7, 9}
	if !findNumber(arr1, 5) {
		t.Errorf("findNumber(arr1, 5) failed")
	}
	if findNumber(arr1, 6) {
		t.Errorf("findNumber(arr1, 6) failed")
	}

	// 测试findLeft函数
	arr2 := []int{1, 2, 3, 4, 4, 4, 5, 6, 7}
	if findLeft(arr2, 4) != 3 {
		t.Errorf("findLeft(arr2, 4) failed")
	}
	if findLeft(arr2, 8) != -1 {
		t.Errorf("findLeft(arr2, 8) failed")
	}

	// 测试findRight函数
	arr3 := []int{1, 2, 3, 4, 4, 4, 5, 6, 7}
	if findRight(arr3, 4) != 5 {
		t.Errorf("findRight(arr3, 4) failed")
	}
	if findRight(arr3, 0) != -1 {
		t.Errorf("findRight(arr3, 0) failed")
	}
}

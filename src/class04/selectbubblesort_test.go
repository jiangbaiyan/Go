package class04

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	N := 200
	V := 1000
	testTimes := 50000
	fmt.Println("测试开始")
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < testTimes; i++ {
		n := rand.Intn(N)
		arr := randomArray(n, V)
		arr1 := copyArray(arr)
		arr2 := copyArray(arr)
		arr3 := copyArray(arr)
		selectSort(arr1)
		bubbleSort(arr2)
		insertSort(arr3)
		if !sameArray(arr1, arr2) || !sameArray(arr1, arr3) {
			fmt.Println("出错了!")
		}
	}
	fmt.Println("测试结束")
}

func randomArray(n, v int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(v) + 1
	}
	return arr
}

func copyArray(arr []int) []int {
	n := len(arr)
	ans := make([]int, n)
	copy(ans, arr)
	return ans
}

func sameArray(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

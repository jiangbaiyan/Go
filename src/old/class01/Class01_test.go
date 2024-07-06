package class01

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{6, 3, 4, 2, 1}
	SelectionSort(arr)
	specArr := []int{1, 2, 3, 4, 6}
	assert.Equal(t, true, reflect.DeepEqual(arr, specArr))
	t.Log(arr)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{6, 3, 4, 2, 1}
	BubbleSort(arr)
	specArr := []int{1, 2, 3, 4, 6}
	assert.Equal(t, true, reflect.DeepEqual(arr, specArr))
	t.Log(arr)
}

func TestInsertionSort(t *testing.T) {
	arr := []int{6, 3, 4, 2, 1}
	InsertionSort(arr)
	specArr := []int{1, 2, 3, 4, 6}
	assert.Equal(t, true, reflect.DeepEqual(arr, specArr))
	t.Log(arr)
}

func TestBSExist(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 6, 7, 8, 10}
	assert.Equal(t, false, BSExist(arr, 5))
	assert.Equal(t, true, BSExist(arr, 2))
	assert.Equal(t, true, BSExist(arr, 6))
}

func TestBSNearLeft(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 6, 7, 8, 10}
	assert.Equal(t, 1, BSNearLeft(arr, 2))
	assert.Equal(t, 3, BSNearLeft(arr, 3))
	assert.Equal(t, 5, BSNearLeft(arr, 5))
}

func TestBSNearRight(t *testing.T) {
	arr := []int{1, 2, 2, 3, 4, 6, 7, 8, 10}
	assert.Equal(t, 2, BSNearRight(arr, 2))
	assert.Equal(t, 3, BSNearRight(arr, 3))
	assert.Equal(t, 4, BSNearRight(arr, 5))
}

func TestBSAwesome(t *testing.T) {
	arr := []int{5, 4, 3, 4, 3, 1, 2}
	assert.Equal(t, 3, BSAwesome(arr))
}

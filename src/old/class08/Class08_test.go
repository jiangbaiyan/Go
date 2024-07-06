package class08

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	arr := []int{6, 3, 4, 2, 1}
	CountSort(arr)
	specArr := []int{1, 2, 3, 4, 6}
	assert.Equal(t, true, reflect.DeepEqual(arr, specArr))
	t.Log(arr)
}

func TestRadixSort(t *testing.T) {
	arr := []int{101, 20, 3, 43, 6}
	RadixSort(arr)
	specArr := []int{3, 6, 20, 43, 101}
	assert.Equal(t, true, reflect.DeepEqual(arr, specArr))
	t.Log(arr)
}

func TestGetDigit(t *testing.T) {
	t.Log(GetDigit(103, 1))
}

package class02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	arr := []int{1, 2, 3, 2, 4}
	Swap(arr, 1, 2)
	assert.Equal(t, 3, arr[1])
	assert.Equal(t, 2, arr[2])
}

func TestOddTimesNum1(t *testing.T) {
	arr := []int{1, 2, 1, 2, 4}
	assert.Equal(t, 4, OddTimesNum1(arr))
}

func TestOddTimesNum2(t *testing.T) {
	arr := []int{1, 2, 1, 2, 4, 5}
	num1, num2 := OddTimesNum2(arr)
	assert.Equal(t, 4, num1)
	assert.Equal(t, 5, num2)
}

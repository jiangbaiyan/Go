package class05

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{6, 3, 4, 2, 1}
	QuickSort(arr)
	specArr := []int{1, 2, 3, 4, 6}
	assert.Equal(t, true, reflect.DeepEqual(arr, specArr))
	t.Log(arr)
}

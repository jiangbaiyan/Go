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

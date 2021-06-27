package class08

import (
	"github.com/jiangbaiyan/go/src/class03"
	"math"
)

// RadixSort 基数排序, 适用于非负的十进制数
// 1. 找到最高的位数是多少, 最高位填充0
// 2. 个位数进桶出桶
// 3. 十位数进桶出桶
// 4. ...最高位进桶出桶, 最终有序
func RadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// 1. 准备10个桶
	buckets := make([]*class03.MyRingQueue, 10, 10)
	digit := MaxBits(arr)
	// d = 1为个位, 从最低位一直到最高位
	for d := 1; d <= digit; d++ {
		for _, item := range arr {
			if buckets[GetDigit(item, d)] == nil {
				buckets[GetDigit(item, d)] = class03.NewMyRingQueue(10000)
			}
			buckets[GetDigit(item, d)].Push(item)
		}
		i := 0
		for _, item := range buckets {
			if item != nil {
				for !item.IsEmpty() {
					arr[i] = item.Pop().(int)
					i++
				}
			}
		}
	}
}

func GetDigit(x int, d int) int {
	return (x / (int(math.Pow10(d - 1)))) % 10
}

// MaxBits 求该数有几位
func MaxBits(arr []int) int {
	max := -99999999
	for _, item := range arr {
		max = getMax(max, item)
	}
	res := 0
	for max != 0 {
		max /= 10
		res++
	}
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

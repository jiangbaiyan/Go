package class08

// RadixSort 基数排序, 适用于非负的十进制数
// 1. 找到最高的位数是多少, 最高位填充0
// 2. 个位数进桶出桶
// 3. 十位数进桶出桶
// 4. ...最高位进桶出桶, 最终有序
func RadixSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

}

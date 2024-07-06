package class02

// Swap 不用临时变量交换两个数
// 当然go语言可以直接搞
func Swap(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

package class02

// arr中，只有一种数，出现奇数次
func OddTimesNum1(arr []int) int {
	eor := 0
	for _, item := range arr {
		eor ^= item
	}
	return eor
}

// arr中，有两种数，出现奇数次
func OddTimesNum2(arr []int) (int, int) {
	// 1. 先拿到这两个数的疑惑结果，用eor表示，假设这两个数为4和5
	eor := 0
	for _, item := range arr {
		eor ^= item
	}
	// eor = 4 ^ 5  eor != 0
	// 2. 因为要分离4和5，必须要找到这俩的不同点. 4和5的不同点就在于某一个人某一位上一定是1，另一个为0。所以要先找到eor最右边的1
	rightOne := eor & -eor // rightOne = 000000001000
	eorPie := 0
	for _, item := range arr {
		// 和rightOne做与运算，若为0则说明item那一位上为0, 把item为0的都异或起来，最后剩的就是其中一个，这里得到的就是4
		if item&rightOne == 0 {
			eorPie ^= item
		}
	}
	// eor = 4 ^ 5; eorPie = 4; eor ^ eorPie = 4 ^ 5 ^ 4 = 5
	return eorPie, eor ^ eorPie
}

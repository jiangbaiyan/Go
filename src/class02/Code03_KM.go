package class02

// OnlyKTimes 一个数组中有一种数出现K次，其他数都出现了M次，
// M > 1,  K < M
// 找到，出现了K次的数，
// 要求，额外空间复杂度O(1)，时间复杂度O(N)
func OnlyKTimes(arr []int, k int, m int) int {
	// 构建统计1的切片
	t := make([]int, 32, 32)
	for _, item := range arr {
		for i := 0; i < 32; i++ {
			if (item>>i)&1 == 1 {
				t[i]++
			}
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		// 如果这一位上能整除m，那么k在这一位上为0，不管
		if t[i]%m == 0 {
			continue
		} else if t[i]%m == k {
			ans |= 1 << i
		} else {
			return -1
		}
	}
	// ans = 0特殊处理, 因为0 % m 永远等于0
	if ans == 0 {
		count := 0
		for _, item := range arr {
			if item == 0 {
				count++
			}
		}
		if count != k {
			return -1
		}
	}
	return ans
}

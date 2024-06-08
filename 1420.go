package main

// 1 <= val <= m
// k

func numOfArrays(n int, m int, k int) int {
	if k == 0 {
		return 0
	}
	cache := map[[3]int]int{}
	var helper func(leftSize int, val int, leftK int) int
	helper = func(leftSize int, val int, leftK int) (ret int) {
		if leftK > leftSize {
			return 0
		}
		if leftK < 0 {
			return 0
		}
		if c, ok := cache[[3]int{leftSize, val, leftK}]; ok {
			return c
		}
		defer func() {
			cache[[3]int{leftSize, val, leftK}] = ret
		}()
		if leftSize == 0 {
			return 1
		}
		for next := 1; next <= val; next++ {
			ret += helper(leftSize-1, val, leftK)
		}
		for next := val + 1; next <= m; next++ {
			ret += helper(leftSize-1, next, leftK-1)
		}
		ret %= 1e9 + 7
		return ret
	}
	return helper(n, 0, k)
}

package main

import "math/bits"

func findIntegers(n int) int {
	size := bits.Len(uint(n))
	dp := make([]int, size+1)
	for i := 2; i <= size; i++ {
		dp[i] = dp[i-1] + dp[i-2] + 1<<(i-2)
	}
	ret := n + 1
	for i := size - 1; n > 0 && i >= 0; i-- {
		if (1<<i)&n == 0 {
			continue
		}
		ret -= dp[i]
		n -= 1 << i
		if n >= 1<<(i-1) {
			n -= 1 << (i - 1)
			ret -= n - (1 << (i - 1)) + 1
		}
		i -= 1
	}
	return ret
}

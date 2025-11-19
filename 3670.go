package main

import "math/bits"

func maxProduct9(nums []int) int64 {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	size := bits.Len(uint(maxVal))
	dp := make([]int, 1<<size)
	for _, num := range nums {
		dp[num] = num
	}
	for i := 0; i < len(dp); i++ {
		if dp[i] == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if i&(1<<j) == 0 {
				continue
			}
			dp[i] = max(dp[i], dp[i^(1<<j)]*dp[1<<j])
		}
	}
	mask := 1<<size - 1
	var ret int
	for _, num := range nums {
		rev := mask ^ num
		ret = max(ret, num*dp[rev])
	}
	return int64(ret)
}

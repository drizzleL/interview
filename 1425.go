package main

import (
	"math"
)

func constrainedSubsetSum(nums []int, k int) int {
	dp := make([]int, len(nums))
	var stack []int
	ret := math.MinInt32
	for i, num := range nums {
		var tmp int
		if len(stack) != 0 {
			tmp = max(tmp, stack[0])
		}
		curr := num + tmp
		dp[i] = curr
		ret = max(ret, curr)
		if i >= k && len(stack) > 0 && stack[0] == dp[i-k] {
			stack = stack[1:]
		}
		for len(stack) > 0 && stack[len(stack)-1] < curr {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, curr)
	}
	return ret
}

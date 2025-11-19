package main

import "math"

func maximumScore4(nums []int, k int) int64 {
	helper := func(nums []int) int {
		dp := make([][]int, k*2+1)
		for j := range dp {
			dp[j] = make([]int, 3)
			for m := range dp[j] {
				dp[j][m] = math.MinInt32
			}
		}
		for i, num := range nums {
			if i == 0 {
				dp[0][1] = 0
				dp[1][2] = num
				dp[1][0] = -num
				continue
			}
			for j := k * 2; j >= 1; j-- {
				for m := 1; m <= 2; m++ {
					dp[j][m] = max(dp[j][m], dp[j-1][m-1]+num)
				}
				for m := 0; m <= 1; m++ {
					dp[j][m] = max(dp[j][m], dp[j-1][m+1]-num)
				}
			}
		}
		var ret int
		for j := 2; j <= k*2; j += 2 {
			ret = max(ret, dp[j][1])
		}
		return ret
	}
	var minIdx int
	for i, num := range nums {
		if num < nums[minIdx] {
			minIdx = i
		}
	}
	size := len(nums)
	nums = append(nums, nums...)
	ret := helper(nums[minIdx : minIdx+size])
	nums = nums[minIdx+1 : minIdx+size+1]
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return int64(max(ret, helper(nums)))
}

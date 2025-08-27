package main

import "math"

func minXor2(nums []int, k int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, k)
	}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = nums[i]
		if i != 0 {
			dp[i][0] ^= dp[i-1][0]
		}
	}
	helper := func(i, j int) int {
		ret := dp[j][0]
		if i != 0 {
			ret ^= dp[i-1][0]
		}
		return ret
	}
	for j := 1; j < k; j++ {
		for i := len(nums) - 1; i >= 0; i-- {
			dp[i][j] = math.MaxInt32
			for mid := i; mid >= 0; mid-- {
				if mid < j {
					break
				}
				dp[i][j] = min(dp[i][j], max(helper(mid, i), dp[mid-1][j-1]))
			}
		}
	}
	return dp[len(nums)-1][k-1]
}
func minXor(nums []int, k int) int {
	presum := make([]int, len(nums)+1)
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] ^ nums[i-1]
	}
	cache := map[[3]int]int{}
	var helper func(i, j int, k int) int
	helper = func(i, j int, k int) (ret int) {
		if k == 0 {
			return 0
		}
		if k == 1 {
			return presum[j+1] ^ presum[i]
		}
		if c, ok := cache[[3]int{i, j, k}]; ok {
			return c
		}
		defer func() {
			cache[[3]int{i, j, k}] = ret
		}()
		ret = math.MaxInt32
		for m := i; m < j; m++ {
			if j-m < k-1 {
				break
			}
			ret = min(ret, max(helper(i, m, 1), helper(m+1, j, k-1)))
		}
		return ret
	}
	return helper(0, len(nums)-1, k)
}

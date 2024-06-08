package main

import "math"

func minSpaceWastedKResizing(nums []int, k int) int {
	var helper func(i, j int, k int) int
	dp := make([][][]int, len(nums))
	for i := range dp {
		dp[i] = make([][]int, len(nums))
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	helper = func(i, j int, k int) (ret int) {
		if dp[i][j][k] != -1 {
			return dp[i][j][k]
		}
		defer func() {
			dp[i][j][k] = ret
		}()
		if j-i+1 <= k {
			return 0
		}
		if k == 0 {
			var sum int
			var peak int
			for m := i; m <= j; m++ {
				sum += nums[m]
				peak = max(peak, nums[m])
			}
			return peak*(j-i+1) - sum
		}
		tmp := math.MaxInt32
		for m := i; m < j; m++ {
			tmp = min(tmp, helper(i, m, 0)+helper(m+1, j, k-1))
		}
		return tmp
	}
	return helper(0, len(nums)-1, k)
}

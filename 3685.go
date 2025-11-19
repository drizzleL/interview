package main

import "sort"

func subsequenceSumAfterCapping(nums []int, k int) []bool {
	sort.Ints(nums)
	ret := make([]bool, len(nums))
	dp := make([]bool, k+1)
	for i, j := 0, 0; i < len(nums); i++ {
		x := i + 1
		for j < len(nums) && nums[j] <= x {
			for m := k; m >= 0; m-- {
				dp[m] = dp[m] || dp[m-nums[j]]
			}
			j++
		}
		leftSize := len(nums) - j
		for m := 0; m <= leftSize; m++ {
			if dp[k-m*x] {
				ret[i] = true
				break
			}
		}
	}
	return ret
}

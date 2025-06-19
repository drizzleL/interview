package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	var ret []int
	for i := 0; i < len(nums); i++ {
		var tmp []int
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(dp[j]) > len(tmp) {
				tmp = dp[j]
			}
		}
		tmpCpy := make([]int, len(tmp))
		copy(tmpCpy, tmp)
		dp[i] = append(tmpCpy, nums[i])
		if len(dp[i]) > len(ret) {
			ret = dp[i]
		}
	}
	return ret
}

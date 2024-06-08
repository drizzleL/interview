package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([][2]int, len(nums))
	var idx int
	for i := 0; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] != 0 {
				continue
			}
			if dp[i][1] < dp[j][1]+1 {
				dp[i] = [2]int{j, dp[j][1] + 1}
			}
		}
		if dp[i][1] == 0 {
			dp[i] = [2]int{i, 1}
		}
		if dp[i][1] > dp[idx][1] {
			idx = i
		}
	}
	var ret []int
	for dp[idx][0] != idx {
		ret = append(ret, nums[idx])
		idx = dp[idx][0]
	}
	ret = append(ret, nums[idx])
	for i, j := 0, len(ret)-1; i < j; i++ {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

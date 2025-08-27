package main

import "sort"

func sumOfPowers(nums []int, k int) int {
	sort.Ints(nums)
	diff := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			diff[nums[j]-nums[i]] = true
		}
	}
	var helper func(d int) int
	helper = func(d int) int {
		dp := make([][][]int, len(nums))
		for i := range dp {
			dp[i] = make([][]int, k)
			for j := range dp[i] {
				dp[i][j] = make([]int, 2)
			}
			dp[i][0][0] = 1
		}
		for k2 := 1; k2 < k; k2++ {
			for i := 0; i < len(nums); i++ {
				for j := 0; j < i; j++ {
					if nums[i]-nums[j] < d {
						break
					}
					dp[i][k2][1] += dp[j][k2-1][1]
					dp[i][k2][1] %= 1e9 + 7
					if nums[i]-nums[j] == d {
						dp[i][k2][1] += dp[j][k2-1][0]
						dp[i][k2][1] %= 1e9 + 7
					} else {
						dp[i][k2][0] += dp[j][k2-1][0]
						dp[i][k2][0] %= 1e9 + 7
					}
				}
			}
		}
		var ret int
		for i := 0; i < len(nums); i++ {
			ret += dp[i][k-1][1]
			ret %= 1e9 + 7
		}
		return ret
	}
	var ret int
	for d := range diff {
		ret += helper(d) * d
		ret %= 1e9 + 7
	}
	return ret
}

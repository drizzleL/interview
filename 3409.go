package main

func longestSubsequence3(nums []int) int {
	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}
	maxDiff := maxVal - minVal
	dp := make([][]int, maxVal+1)
	for i := range dp {
		dp[i] = make([]int, maxDiff+1)
	}
	var ret int
	for _, v := range nums {
		for prev := minVal; prev <= maxVal; prev++ {
			diff := abs(prev - v)
			dp[v][diff] = max(dp[v][diff], dp[prev][diff]+1)
			ret = max(ret, dp[v][diff])
		}
		for diff := maxVal - minVal - 1; diff >= 0; diff-- {
			dp[v][diff] = max(dp[v][diff], dp[v][diff+1])
		}
	}
	return ret
}

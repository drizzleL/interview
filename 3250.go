package main

func countOfPairs(nums []int) int {
	var maxVal int
	for _, num := range nums {
		maxVal = max(maxVal, num)
	}
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, maxVal+1)
	}
	for i := nums[0]; i >= 0; i-- {
		dp[0][i] = 1
	}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		preNum := nums[i-1]
		for j := num; j >= 0; j-- {
			nowB := num - j
			for preA := j; preA >= 0; preA-- {
				preB := preNum - preA
				if preB < 0 || nowB > preB { // can't reach
					continue
				}
				dp[i][j] += dp[i-1][preA]
				dp[i][j] %= 1e9 + 7
			}
		}
	}
	var ret int
	for i := 0; i <= maxVal; i++ {
		ret += dp[len(nums)-1][i]
		ret %= 1e9 + 7
	}
	return ret
}

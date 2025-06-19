package main

func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum < target || target < -sum || (target+sum)%2 != 0 {
		return 0
	}
	dp := make([]int, (sum+target)/2+1)
	dp[0] = 1
	for _, num := range nums {
		for j := len(dp) - 1; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[len(dp)-1]
}

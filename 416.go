package main

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for j := len(dp) - 1; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[sum]
}

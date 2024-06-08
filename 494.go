package main

func findTargetSumWays(nums []int, target int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	lastdp := make([]int, sum*2+1)
	lastdp[sum] = 1
	for _, num := range nums {
		dp := make([]int, sum*2+1)
		for i := range lastdp {
			if i-num >= 0 {
				dp[i] += lastdp[i-num]
			}
			if i+num < len(lastdp) {
				dp[i] += lastdp[i+num]
			}
		}
		lastdp = dp
	}
	return lastdp[target+sum]
}

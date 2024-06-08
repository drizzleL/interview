package main

func lastStoneWeightII(stones []int) int {
	var sum int
	for _, num := range stones {
		sum += num
	}
	dp := make([]bool, sum/2+1)
	dp[0] = true
	for _, stone := range stones {
		for i := len(dp) - 1; i >= stone; i-- {
			dp[i] = dp[i] || dp[i-stone]
		}
	}
	sumN := sum / 2
	for ; sumN >= 0; sumN-- {
		if dp[sumN] {
			break
		}
	}
	return sum - 2*sumN
}

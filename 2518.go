package main

func countPartitions2(nums []int, k int) int {
	ret := fastPow(2, len(nums), 1)
	if ret < 0 {
		ret += 1e9 + 7
	}
	dp := make([]int, k)
	dp[0] = 1
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		for j := k - 1; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
			dp[j] %= 1e9 + 7
		}
	}
	if total < k*2 {
		return 0
	}
	var sum int
	for i := 0; i < k; i++ {
		sum += dp[i]
		sum %= 1e9 + 7
	}
	sum *= 2
	sum %= 1e9 + 7
	ret -= sum
	if ret < 0 {
		ret += 1e9 + 7
	}
	return ret
}

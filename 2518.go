package main

import "log"

func countPartitions2(nums []int, k int) int {
	ret := fastPow(2, len(nums), 1)
	ret -= 1
	if ret < 0 {
		ret += 1e9 + 7
	}
	dp := make([]int, k)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := k - 1; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
			dp[j] %= 1e9 + 7
		}
	}
	var sum int
	for i := 1; i < k; i++ {
		sum += dp[i]
		sum %= 1e9 + 7
	}
	log.Println(ret, sum)
	sum *= 2
	sum %= 1e9 + 7
	ret -= sum
	if ret < 0 {
		ret += 1e9 + 7
	}
	return ret
}

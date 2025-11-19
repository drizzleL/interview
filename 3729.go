package main

func numGoodSubarrays(nums []int, k int) int64 {
	var presum int
	dp := make([]int, k)
	dp[0] = 1
	var ret int
	for _, num := range nums {
		presum += num
		rem := presum % k
		ret += dp[rem]
		ret %= 1e9 + 7
		dp[rem] += 1
		dp[rem] %= 1e9 + 7
	}
	for i := 0; i < len(nums); {
		val := nums[i]
		var cnt int
		j := i
		for j < len(nums) && nums[j] == val {
			cnt += 1
			j += 1
		}
		var sum int
		for i < j {
			sum += nums[i]
			if sum%k == 0 {
				ret -= j - i - 1
				if ret < 0 {
					ret += 1e9 + 7
				}
			}
			i += 1
		}
	}
	return int64(ret)
}

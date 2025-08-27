package main

func minArraySum3(nums []int, k int) int64 {
	var sum int
	dp := make(map[int]int, k) // save best idx + 1
	dp[0] = 0
	maxs := make([]int, len(nums)+1)
	presums := make([]int, len(nums)+1)
	for i, num := range nums {
		maxs[i+1] = maxs[i]
		presums[i+1] = presums[i] + num
		sum += num
		rem := sum % k

		old, ok := dp[rem]
		if !ok {
			dp[rem] = i + 1
			continue
		}
		maxs[i+1] = max(maxs[i+1], presums[i+1]-presums[old]+maxs[old])
		if maxs[i+1]-presums[i] > maxs[old]-presums[old] {
			dp[rem] = i + 1
		}
	}
	return int64(sum - maxs[len(nums)])
}

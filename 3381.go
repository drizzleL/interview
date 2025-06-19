package main

func maxSubarraySum(nums []int, k int) int64 {
	sums := make([]int, k)
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
		sums[i] = sum
	}
	sums[k-1] = min(0, sums[k-1])
	ret := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i]
		ret = max(ret, sum-sums[i%k])
		sums[i%k] = min(sums[i%k], sum)
	}
	return int64(ret)
}

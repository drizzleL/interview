package main

import "sort"

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	var ret int
	mid := len(nums) / 2
	ret += abs(nums[mid] - k)
	for i := mid - 1; i >= 0 && nums[i] > k; i-- {
		ret += nums[i] - k
	}
	for i := mid + 1; i < len(nums) && nums[i] < k; i++ {
		ret += k - nums[i]
	}
	return int64(ret)
}

package main

func countNonDecreasingSubarrays(nums []int, k int) int64 {
	ret := 1
	for i, r := len(nums)-2, len(nums)-1; i >= 0; i-- {
		if nums[i] > nums[i+1] {

			continue
		}
		ret += r - i + 1

	}
	return int64(ret)
}

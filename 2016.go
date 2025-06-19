package main

func maximumDifference(nums []int) int {
	minVal := nums[0]
	ret := -1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= minVal {
			minVal = nums[i]
		}
		ret = max(ret, nums[i]-minVal)
	}
	return ret
}

package main

func maximumTripletValue2(nums []int) int64 {
	var ret int
	maxGap := nums[0] - nums[1]
	maxVal := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		ret = max(ret, maxGap*nums[i])
		maxGap = max(maxGap, maxVal-nums[i])
		maxVal = max(maxVal, nums[i])
	}
	return int64(ret)
}

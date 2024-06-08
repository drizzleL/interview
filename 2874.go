package main

func maximumTripletValue(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	var ret int
	maxVals := make([]int, len(nums))
	maxVals[len(maxVals)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		maxVals[i] = max(maxVals[i+1], nums[i])
	}
	maxVal := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		ret = max(ret, (maxVal-nums[i])*maxVals[i+1])
		maxVal = max(maxVal, nums[i])
	}
	return int64(ret)
}

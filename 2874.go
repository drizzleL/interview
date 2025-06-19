package main

func maximumTripletValue(nums []int) int64 {
	var maxGap, minVal, maxVal int
	var ret int
	for i := 0; i < len(nums); i++ {
		ret = max(ret, nums[i]*maxGap)
		if nums[i] > maxVal {
			maxVal = nums[i]
			minVal = nums[i]
			continue
		}
		minVal = min(minVal, nums[i])
		maxGap = max(maxGap, maxVal-minVal)
	}
	return int64(ret)
}

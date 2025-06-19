package main

import "math"

func maximumProduct(nums []int, m int) int64 {
	if m == 1 {
		var ret int
		for _, num := range nums {
			ret = max(ret, num*num)
		}
		return int64(ret)
	}
	ret := math.MinInt64
	minVal, maxVal := nums[len(nums)-1], nums[len(nums)-1]
	for i := len(nums) - m; i >= 0; i-- {
		ret = max(ret, nums[i]*minVal)
		ret = max(ret, nums[i]*maxVal)
		minVal = min(minVal, nums[i+m-2])
		maxVal = max(maxVal, nums[i+m-2])
	}
	return int64(ret)
}

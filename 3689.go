package main

func maxTotalValue(nums []int, k int) int64 {
	maxVal, minVal := nums[0], nums[0]
	for _, num := range nums {
		maxVal = max(maxVal, num)
		minVal = min(minVal, num)
	}
	return int64((maxVal - minVal) * k)
}

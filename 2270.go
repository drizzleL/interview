package main

func waysToSplitArray(nums []int) int {
	var leftSum, rightSum, ret int
	for _, num := range nums {
		rightSum += num
	}
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]
		rightSum -= nums[i]
		if leftSum >= rightSum {
			ret += 1
		}
	}
	return ret
}

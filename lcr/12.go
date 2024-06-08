package main

func pivotIndex(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	var presum int
	for i := range nums {
		if presum*2+nums[i] == sum {
			return i
		}
		presum += nums[i]
	}
	return -1
}

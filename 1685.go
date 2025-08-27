package main

func getSumAbsoluteDifferences(nums []int) []int {
	presums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		presums[i] = presums[i-1] + nums[i-1]
	}
	ret := make([]int, len(nums))
	for i := range nums {
		ret[i] += nums[i]*i - presums[i]
		ret[i] += presums[len(nums)] - presums[i+1] - nums[i]*(len(nums)-i-1)
	}
	return ret
}

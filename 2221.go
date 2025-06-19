package main

func triangularSum(nums []int) int {
	for len(nums) > 1 {
		for j := 0; j < len(nums)-1; j++ {
			nums[j] += nums[j+1] % 10
			nums[j] %= 10
		}
		nums = nums[:len(nums)-1]
	}
	return nums[0]
}

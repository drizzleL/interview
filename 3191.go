package main

func minOperations11(nums []int) int {
	var ret int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] == 1 {
			continue
		}
		ret += 1
		nums[i+1] ^= 1
		nums[i+2] ^= 1
	}
	for i := len(nums) - 2; i < len(nums); i++ {
		if nums[i] == 1 {
			continue
		}
		return -1
	}
	return ret
}

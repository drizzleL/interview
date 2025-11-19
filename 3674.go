package main

func minOperations71(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			return 1
		}
	}
	return 0
}

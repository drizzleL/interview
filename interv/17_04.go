package main

func missingNumber(nums []int) int {
	for i := range nums {
		for i != nums[i] && nums[i] != len(nums) {
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

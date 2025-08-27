package main

func isTrionic(nums []int) bool {
	if len(nums) <= 3 {
		return false
	}
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		break
	}
	if i == len(nums) {
		return false
	}
	if nums[i] == nums[i-1] {
		return false
	}
	for ; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			continue
		}
		break
	}
	if i == len(nums) {
		return false
	}
	if nums[i] == nums[i-1] {
		return false
	}
	for ; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		break
	}
	return i == len(nums)
}

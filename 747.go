package main

func dominantIndex(nums []int) int {
	large1, large2 := 0, 1
	if nums[large2] > nums[large1] {
		large1, large2 = large2, large1
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] >= nums[large1] {
			large1, large2 = i, large1
			continue
		}
		if nums[i] >= nums[large2] {
			large2 = i
		}
	}
	if nums[large1] >= nums[large2] {
		return large1
	}
	return -1
}

package main

func pivotArray(nums []int, pivot int) []int {
	ret := make([]int, len(nums))
	for i, j := 0, 0; i < len(nums); i++ {
		ret[i] = pivot
		if nums[i] < pivot {
			ret[j] = nums[i]
			j++
		}
	}
	for i, j := len(nums)-1, len(nums)-1; i >= 0; i-- {
		if nums[i] > pivot {
			ret[j] = nums[i]
			j--
		}
	}
	return ret
}
